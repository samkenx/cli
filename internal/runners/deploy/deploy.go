package deploy

import (
	"errors"
	"os"
	"strings"

	"github.com/ActiveState/cli/internal/failures"
	"github.com/ActiveState/cli/internal/locale"
	"github.com/ActiveState/cli/internal/logging"
	"github.com/ActiveState/cli/internal/output"
	"github.com/ActiveState/cli/internal/subshell"
	"github.com/ActiveState/cli/internal/virtualenvironment"
	"github.com/ActiveState/cli/pkg/platform/model"
	"github.com/ActiveState/cli/pkg/platform/runtime"
	"github.com/ActiveState/cli/pkg/project"
)

var (
	FailNoCommitForProject = failures.Type("deploy.fail.nocommit")
)

type Params struct {
	Namespace project.Namespaced
	Path      string
	Step      Step
}

type Deploy struct {
	output output.Outputer

	DefaultBranchForProjectName defaultBranchForProjectNameFunc
	NewRuntimeInstaller         newInstallerFunc
}

func NewDeploy(out output.Outputer) *Deploy {
	return &Deploy{
		out,
		model.DefaultBranchForProjectName,
		newInstaller,
	}
}

func (d *Deploy) Run(params *Params) error {
	installer, err := d.createInstaller(params.Namespace, params.Path)
	if err != nil {
		return err
	}

	return runSteps(installer, params.Step, d.output)
}

func (d *Deploy) createInstaller(namespace project.Namespaced, path string) (installable, *failures.Failure) {
	branch, fail := d.DefaultBranchForProjectName(namespace.Owner, namespace.Project)
	if fail != nil {
		return nil, fail
	}

	if branch.CommitID == nil {
		return nil, FailNoCommitForProject.New(locale.Tr("err_deploy_no_commits", namespace.String()))
	}

	return d.NewRuntimeInstaller(*branch.CommitID, namespace.Owner, namespace.Project, path)
}

func runSteps(installer installable, step Step, out output.Outputer) error {
	return runStepsWithFuncs(installer, step, out, install, configure, report)
}

func runStepsWithFuncs(installer installable, step Step, out output.Outputer, installf installFunc, configuref configureFunc, reportf reportFunc) error {
	logging.Debug("runSteps: %s", step.String())

	var envGetter runtime.EnvGetter
	var fail *failures.Failure

	if step == UnsetStep || step == InstallStep {
		logging.Debug("Running install step")
		var err error
		if envGetter, err = installf(installer, out); err != nil {
			return err
		}

		if step == UnsetStep {
			out.Notice("") // Some space between steps
		}
	}
	if step == UnsetStep || step == ConfigureStep {
		logging.Debug("Running configure step")
		if envGetter == nil {
			if envGetter, fail = installer.Env(); fail != nil {
				return fail
			}
		}
		if err := configuref(envGetter, out); err != nil {
			return err
		}
		if step == UnsetStep {
			out.Notice("") // Some space between steps
		}
	}
	if step == UnsetStep || step == ReportStep {
		logging.Debug("Running report step")
		if envGetter == nil {
			if envGetter, fail = installer.Env(); fail != nil {
				return fail
			}
		}
		if err := reportf(envGetter, out); err != nil {
			return err
		}
	}

	return nil
}

type installFunc func(installer installable, out output.Outputer) (runtime.EnvGetter, error)

func install(installer installable, out output.Outputer) (runtime.EnvGetter, error) {
	out.Notice(locale.T("deploy_install"))
	envGetter, installed, fail := installer.Install()
	if fail != nil {
		return envGetter, fail.ToError()
	}
	if ! installed {
		out.Notice(locale.T("using_cached_env"))
	}
	return envGetter, nil
}

type configureFunc func(envGetter runtime.EnvGetter, out output.Outputer) error

func configure(envGetter runtime.EnvGetter, out output.Outputer) error {
	sshell, fail := subshell.Get()
	if fail != nil {
		return fail.ToError()
	}

	venv := virtualenvironment.New(envGetter.GetEnv)
	env := venv.GetEnv(false, "")

	if len(env) == 0 {
		return errors.New(locale.T("err_deploy_run_install"))
	}

	out.Notice(locale.Tr("deploy_configure_shell", sshell.Shell()))

	return sshell.WriteUserEnv(env).ToError()
}

type Report struct {
	BinaryDirectories []string
	Environment       map[string]string
}

type reportFunc func(envGetter runtime.EnvGetter, out output.Outputer) error

func report(envGetter runtime.EnvGetter, out output.Outputer) error {
	venv := virtualenvironment.New(envGetter.GetEnv)
	env := venv.GetEnv(false, "")

	if len(env) == 0 {
		return errors.New(locale.T("err_deploy_run_install"))
	}

	var bins []string

	if path, ok := env["PATH"]; ok {
		delete(env, "PATH")
		bins = strings.Split(path, string(os.PathListSeparator))
	}

	out.Notice(locale.T("deploy_info"))

	out.Print(Report{
		BinaryDirectories: bins,
		Environment:       env,
	})

	out.Notice(locale.T("deploy_restart_shell"))

	return nil
}
