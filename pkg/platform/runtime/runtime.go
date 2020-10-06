package runtime

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/go-openapi/strfmt"

	"github.com/ActiveState/cli/internal/config"
	"github.com/ActiveState/cli/internal/constants"
	"github.com/ActiveState/cli/internal/errs"
	"github.com/ActiveState/cli/internal/failures"
	"github.com/ActiveState/cli/internal/fileutils"
	"github.com/ActiveState/cli/internal/hash"
	"github.com/ActiveState/cli/internal/language"
)

type Runtime struct {
	RuntimeDir  string
	CommitID    strfmt.UUID
	Owner       string
	ProjectName string
}

func NewRuntime(commitID strfmt.UUID, owner string, projectName string) *Runtime {
	return &Runtime{InstallPath(owner, projectName), commitID, owner, projectName}
}

func InstallPath(owner, projectName string) string {
	if runtime.GOOS == "darwin" {
		// mac doesn't use relocation so we can safely use a longer path
		return filepath.Join(config.CachePath(), owner, projectName)
	} else {
		return filepath.Join(config.CachePath(), hash.ShortHash(owner, projectName))
	}
}

// Env will grab the environment information for the given runtime.
// This currently just aliases to installer, pending further refactoring
func (r *Runtime) Env() (EnvGetter, *failures.Failure) {
	return NewInstaller(r).Env()
}

// Languages returns a slice of languages that is supported by the current runtime
func (r *Runtime) Languages() ([]*language.Language, error) {
	installer := NewInstaller(r)
	env, fail := installer.Env()
	if fail != nil {
		return nil, errs.Wrap(fail, "Could not initialize environment information for runtime")
	}

	envMap, err := env.GetEnv(false, "")
	if err != nil {
		return nil, errs.Wrap(err, "Could not get environment information for runtime")
	}

	// Retrieve artifact binary directory
	var bins []string
	if p, ok := envMap["PATH"]; ok {
		bins = strings.Split(p, string(os.PathListSeparator))
	}

	result := []*language.Language{}
	for _, bin := range bins {
		if fileutils.TargetExists(filepath.Join(bin, constants.ActivePython2Executable)) {
			lang := language.Python2
			result = append(result, &lang)
		}
		if fileutils.TargetExists(filepath.Join(bin, constants.ActivePython3Executable)) {
			lang := language.Python3
			result = append(result, &lang)
		}
		if fileutils.TargetExists(filepath.Join(bin, constants.ActivePerlExecutable)) {
			lang := language.Perl
			result = append(result, &lang)
		}
	}

	return result, nil
}