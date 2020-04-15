package integration

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

	"github.com/ActiveState/cli/internal/constants"
	"github.com/ActiveState/cli/internal/fileutils"
	"github.com/ActiveState/cli/internal/locale"
	"github.com/ActiveState/cli/internal/testhelpers/e2e"
)

type ActivateIntegrationTestSuite struct {
	suite.Suite
}

func (suite *ActivateIntegrationTestSuite) TestActivatePython3() {
	suite.activatePython("3")
}

func (suite *ActivateIntegrationTestSuite) TestActivatePython3_zsh() {
	if _, err := exec.LookPath("zsh"); err != nil {
		suite.T().Skip("This test requires a zsh shell in your PATH")
	}
	suite.activatePython("3", "SHELL=zsh")
}

func (suite *ActivateIntegrationTestSuite) TestActivatePython2() {
	suite.activatePython("2")
}

func (suite *ActivateIntegrationTestSuite) TestActivateWithoutRuntime() {
	ts := e2e.New(suite.T(), false)
	defer ts.Close()
	ts.LoginAsPersistentUser()

	cp := ts.Spawn("activate", "ActiveState-CLI/Python3")
	cp.Expect("Where would you like to checkout")
	cp.SendLine(cp.WorkDirectory())
	cp.Expect("activated state", 20*time.Second)
	cp.WaitForInput(10 * time.Second)

	cp.SendLine("exit 123")
	cp.ExpectExitCode(123, 10*time.Second)
}

func (suite *ActivateIntegrationTestSuite) TestActivatePythonByHostOnly() {
	if runtime.GOOS != "linux" {
		suite.T().Skip("not currently testing this OS")
	}

	ts := e2e.New(suite.T(), false)
	defer ts.Close()
	ts.LoginAsPersistentUser()

	projectName := "Python-LinuxWorks"
	cp := ts.Spawn("activate", "cli-integration-tests/"+projectName, "--path="+ts.Dirs.Work)

	cp.Expect("Activating state")
	cp.Expect("activated state", 120*time.Second)
	cp.WaitForInput()
	cp.SendLine("exit")
	cp.ExpectExitCode(0)
}

func (suite *ActivateIntegrationTestSuite) activatePython(version string, extraEnv ...string) {
	// temp skip // pythonExe := "python" + version

	ts := e2e.New(suite.T(), false)
	defer ts.Close()
	ts.LoginAsPersistentUser()

	cp := ts.SpawnWithOpts(
		e2e.WithArgs("activate", "ActiveState-CLI/Python"+version),
		e2e.AppendEnv("ACTIVESTATE_CLI_DISABLE_RUNTIME=false"),
		e2e.AppendEnv(extraEnv...),
	)
	cp.Expect("Where would you like to checkout")
	cp.SendLine(cp.WorkDirectory())
	cp.Expect("Downloading", 20*time.Second)
	cp.Expect("Installing", 120*time.Second)
	cp.Expect("activated state", 120*time.Second)

	// ensure that terminal contains output "Installing x/y" with x, y numbers and x=y
	installingString := regexp.MustCompile(
		"Installing *([0-9]+) */ *([0-9]+)",
	).FindAllStringSubmatch(cp.TrimmedSnapshot(), 1)
	suite.Require().Len(installingString, 1, "no match for Installing x / x in\n%s", cp.TrimmedSnapshot())
	suite.Require().Equalf(
		installingString[0][1], installingString[0][2],
		"expected all artifacts are reported to be installed, got %s", installingString[0][0],
	)

	// ensure that shell is functional
	cp.WaitForInput()

	// test python
	// Temporarily skip these lines until MacOS on Python builds with correct copyright
	// temp skip // cp.SendLine(pythonExe + " -c \"import sys; print(sys.copyright)\"")
	// temp skip // cp.Expect("ActiveState Software Inc.")

	// temp skip // cp.SendLine(pythonExe + " -c \"import pytest; print(pytest.__doc__)\"")
	// temp skip // cp.Expect("unit and functional testing")

	// de-activate shell
	cp.SendLine("exit")
	cp.ExpectExitCode(0)
}

func (suite *ActivateIntegrationTestSuite) TestActivatePython3_Forward() {
	var project string
	if runtime.GOOS == "darwin" {
		project = "Activate-MacOS"
	} else {
		project = "Python3"
	}

	ts := e2e.New(suite.T(), false)
	defer ts.Close()

	contents := strings.TrimSpace(fmt.Sprintf(`
project: "https://platform.activestate.com/ActiveState-CLI/%s"
branch: %s
version: %s
`, project, constants.BranchName, constants.Version))

	ts.PrepareActiveStateYAML(contents)

	fmt.Printf("login \n")
	ts.LoginAsPersistentUser()
	fmt.Printf("logged in \n")

	// Ensure we have the most up to date version of the project before activating
	cp := ts.SpawnWithOpts(
		e2e.WithArgs("pull"),
		e2e.AppendEnv("ACTIVESTATE_CLI_DISABLE_RUNTIME=false"),
	)
	cp.Expect("Your activestate.yaml has been updated to the latest version available")
	cp.Expect("If you have any active instances of this project open in other terminals")
	cp.ExpectExitCode(0)

	c2 := ts.Spawn("activate")
	c2.Expect(fmt.Sprintf("Activating state: ActiveState-CLI/%s", project))

	// not waiting for activation, as we test that part in a different test
	c2.WaitForInput()
	c2.SendLine("exit")
	c2.ExpectExitCode(0)
}

func (suite *ActivateIntegrationTestSuite) TestActivatePerl() {
	perlExe := "perl"
	if runtime.GOOS == "darwin" {
		suite.T().Skip("Perl not supported on macOS")
	}

	ts := e2e.New(suite.T(), false)
	defer ts.Close()
	ts.LoginAsPersistentUser()

	cp := ts.SpawnWithOpts(
		e2e.WithArgs("activate", "ActiveState-CLI/Perl"),
		e2e.AppendEnv("ACTIVESTATE_CLI_DISABLE_RUNTIME=false"),
	)
	cp.Expect("Where would you like to checkout")
	cp.SendLine(cp.WorkDirectory())
	cp.Expect("Downloading", 20*time.Second)
	cp.Expect("Installing", 120*time.Second)
	cp.Expect("activated state", 120*time.Second)

	// ensure that terminal contains output "Installing x/y" with x, y numbers and x=y
	installingString := regexp.MustCompile(
		"Installing *([0-9]+) */ *([0-9]+)",
	).FindAllStringSubmatch(cp.TrimmedSnapshot(), 1)
	suite.Require().Len(installingString, 1, "no match for Installing x / x in\n%s", cp.TrimmedSnapshot())
	suite.Require().Equalf(
		installingString[0][1], installingString[0][2],
		"expected all artifacts are reported to be installed, got %s", installingString[0][0],
	)

	// ensure that shell is functional
	cp.WaitForInput()

	cp.SendLine(perlExe + " -e \"use DBD::Pg\"")
	// Expect no output as an error would be printed if the module wasn't available
	cp.Expect("")
}

func (suite *ActivateIntegrationTestSuite) testOutput(method string) {
	ts := e2e.New(suite.T(), false)
	defer ts.Close()

	ts.LoginAsPersistentUser()
	cp := ts.Spawn("activate", "ActiveState-CLI/Python3", "--output", method)
	cp.Expect("Where would you like to checkout")
	cp.SendLine(cp.WorkDirectory())
	cp.Expect("[activated-JSON]")
	cp.ExpectExitCode(0)
}

func (suite *ActivateIntegrationTestSuite) TestActivate_Subdir() {
	ts := e2e.New(suite.T(), false)
	defer ts.Close()
	fail := fileutils.Mkdir(ts.Dirs.Work, "foo", "bar", "baz")
	suite.Require().NoError(fail.ToError())

	// Create the project file at the root of the temp dir
	content := strings.TrimSpace(fmt.Sprintf(`
project: "https://platform.activestate.com/ActiveState-CLI/Python3"
branch: %s
version: %s
`, constants.BranchName, constants.Version))

	ts.PrepareActiveStateYAML(content)

	// Pull to ensure we have an up to date config file
	cp := ts.Spawn("pull")
	cp.Expect("Your activestate.yaml has been updated to the latest version available")
	cp.Expect("If you have any active instances of this project open in other terminals")
	cp.ExpectExitCode(0)

	// Activate in the subdirectory
	c2 := ts.SpawnWithOpts(
		e2e.WithArgs("activate"),
		e2e.WithWorkDirectory(filepath.Join(ts.Dirs.Work, "foo", "bar", "baz")),
	)
	c2.Expect("Activating state: ActiveState-CLI/Python3")

	c2.WaitForInput()
	c2.SendLine("exit")
	c2.ExpectExitCode(0)

}

func (suite *ActivateIntegrationTestSuite) TestInit_Activation_NoCommitID() {
	ts := e2e.New(suite.T(), false)
	defer ts.Close()

	cp := ts.Spawn("init", namespace, "python3")
	cp.Expect(fmt.Sprintf("Project '%s' has been succesfully initialized", namespace))
	cp.ExpectExitCode(0)
	cp = ts.SpawnWithOpts(
		e2e.WithArgs("activate"),
		e2e.AppendEnv("ACTIVESTATE_CLI_DISABLE_RUNTIME=false"),
	)
	cp.Expect(locale.Tr("installer_err_runtime_no_commits", namespace))
	cp.ExpectExitCode(0)
}

func (suite *ActivateIntegrationTestSuite) TestActivate_JSON() {
	suite.testOutput("json")
}

func TestActivateIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(ActivateIntegrationTestSuite))
}
