package deploy

import (
	"os"
	"reflect"
	"testing"

	"github.com/ActiveState/cli/internal/failures"
	"github.com/ActiveState/cli/internal/output"
	"github.com/ActiveState/cli/internal/testhelpers/outputhelper"
	"github.com/ActiveState/cli/pkg/platform/runtime"
)

type InstallableMock struct{}

func (i *InstallableMock) Install() (envGetter EnvGetter, freshInstallation bool, fail *failures.Failure) {
	return nil, false, nil
}

func (i *InstallableMock) Env() (envGetter EnvGetter, fail *failures.Failure) {
	return nil, nil
}

type EnvGetMock struct {
	callback func(inherit bool, projectDir string) (map[string]string, *failures.Failure)
}

func (e *EnvGetMock) GetEnv(inherit bool, projectDir string) (map[string]string, *failures.Failure) {
	return e.callback(inherit, projectDir)
}

func Test_runStepsWithFuncs(t *testing.T) {
	type args struct {
		installer Installable
		step      Step
	}
	type want struct {
		err           error
		installCalled bool
		configCalled  bool
		reportCalled  bool
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			"Deploy without steps",
			args{
				&InstallableMock{},
				UnsetStep,
			},
			want{
				nil,
				true,
				true,
				true,
			},
		},
		{
			"Deploy with install step",
			args{
				&InstallableMock{},
				InstallStep,
			},
			want{
				nil,
				true,
				false,
				false,
			},
		},
		{
			"Deploy with config step",
			args{
				&InstallableMock{},
				ConfigureStep,
			},
			want{
				nil,
				false,
				true,
				false,
			},
		},
		{
			"Deploy with report step",
			args{
				&InstallableMock{},
				ReportStep,
			},
			want{
				nil,
				false,
				false,
				true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var installCalled bool
			installFunc := func(Installable, output.Outputer) (runtime.EnvGetter, error) {
				installCalled = true
				return nil, nil
			}
			var configCalled bool
			configFunc := func(runtime.EnvGetter, output.Outputer) error {
				configCalled = true
				return nil
			}
			var reportCalled bool
			reportFunc := func(runtime.EnvGetter, output.Outputer) error {
				reportCalled = true
				return nil
			}
			catcher := outputhelper.NewCatcher()
			err := runStepsWithFuncs(tt.args.installer, tt.args.step, catcher.Outputer, installFunc, configFunc, reportFunc)
			if err != tt.want.err {
				t.Errorf("runStepsWithFuncs() error = %v, wantErr %v", err, tt.want.err)
			}
			if installCalled != tt.want.installCalled {
				t.Errorf("runStepsWithFuncs() installCalled = %v, want %v", installCalled, tt.want.installCalled)
			}
			if configCalled != tt.want.configCalled {
				t.Errorf("runStepsWithFuncs() configCalled = %v, want %v", configCalled, tt.want.configCalled)
			}
			if reportCalled != tt.want.reportCalled {
				t.Errorf("runStepsWithFuncs() reportCalled = %v, want %v", reportCalled, tt.want.reportCalled)
			}
		})
	}
}

func Test_report(t *testing.T) {
	type args struct {
		envGetter runtime.EnvGetter
	}
	tests := []struct {
		name       string
		args       args
		wantBinary []string
		wantEnv    map[string]string
		wantErr    error
	}{
		{
			"Report",
			args{
				&EnvGetMock{
					func(inherit bool, projectDir string) (map[string]string, *failures.Failure) {
						return map[string]string{
							"KEY1": "VAL1",
							"KEY2": "VAL2",
							"PATH": "PATH1" + string(os.PathListSeparator) + "PATH2",
						}, nil
					},
				},
			},
			[]string{"PATH1", "PATH2"},
			map[string]string{
				"KEY1": "VAL1",
				"KEY2": "VAL2",
			},
			nil,
		},
	}
	for _, tt := range tests {
		catcher := outputhelper.TypedCatcher{}
		t.Run(tt.name, func(t *testing.T) {
			if err := report(tt.args.envGetter, &catcher); err != tt.wantErr {
				t.Errorf("report() error = %v, wantErr %v", err, tt.wantErr)
				t.FailNow()
			}
			report, ok := catcher.Prints[0].(Report)
			if ! ok {
				t.Errorf("Printed unknown structure, expected Report type. Value: %v", report)
				t.FailNow()
			}

			if ! reflect.DeepEqual(report.Environment, tt.wantEnv) {
				t.Errorf("Expected envs to be the same. Want: %v, got: %v", tt.wantEnv, report.Environment)
			}

			if ! reflect.DeepEqual(report.BinaryDirectories, tt.wantBinary) {
				t.Errorf("Expected bins to be the same. Want: %v, got: %v", tt.wantBinary, report.BinaryDirectories)
			}
		})
	}
}
