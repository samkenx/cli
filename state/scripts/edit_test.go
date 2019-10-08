package scripts

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
	"gopkg.in/yaml.v2"

	"github.com/ActiveState/cli/internal/failures"
	"github.com/ActiveState/cli/internal/fileutils"
	"github.com/ActiveState/cli/internal/locale"
	"github.com/ActiveState/cli/internal/scriptfile"
	"github.com/ActiveState/cli/pkg/project"
	"github.com/ActiveState/cli/pkg/projectfile"
)

type EditTestSuite struct {
	suite.Suite
	projectFile    *projectfile.Project
	project        *project.Project
	scriptFile     *scriptfile.ScriptFile
	originalEditor string
}

func (suite *EditTestSuite) BeforeTest(suiteName, testName string) {
	suite.projectFile = &projectfile.Project{}
	contents := strings.TrimSpace(`
project: "https://platform.activestate.com/EditOrg/EditProject?commitID=00010001-0001-0001-0001-000100010001"
constants:
  - name: HELLO
    value: hello
scripts:
  - name: hello
    value: echo hello
  - name: hello-constant
    value: echo $constants.HELLO
  - name: replace
    value: echo replaced
`)

	tempDir := os.TempDir()
	err := os.Chdir(tempDir)
	suite.Require().NoError(err, "should change directories without issue")

	err = yaml.Unmarshal([]byte(contents), suite.projectFile)
	suite.Require().NoError(err, "unexpected error marshalling yaml")

	suite.projectFile.SetPath(filepath.Join(tempDir, "activestate.yaml"))
	fail := suite.projectFile.Save()
	suite.Require().NoError(err, "should be able to save in temp dir")

	suite.project, fail = project.New(suite.projectFile)
	suite.Require().NoError(fail.ToError(), "unexpected error creating project")

	suite.originalEditor = os.Getenv("EDITOR")
}

func (suite *EditTestSuite) AfterTest(suiteName, testName string) {
	err := os.Remove(suite.projectFile.Path())
	suite.Require().NoError(err, "unexpected error removing project file")

	if suite.scriptFile != nil {
		suite.scriptFile.Clean()
	}

	os.Setenv("EDITOR", suite.originalEditor)
}

func (suite *EditTestSuite) TestCreateScriptFile() {
	script := suite.project.ScriptByName("hello")

	var fail *failures.Failure
	suite.scriptFile, fail = createScriptFile(script)
	suite.Require().NoError(fail.ToError(), "should create file")
}

func (suite *EditTestSuite) TestCreateScriptFile_Expand() {
	script := suite.project.ScriptByName("hello-constant")

	EditFlags.Expand = true
	var fail *failures.Failure
	suite.scriptFile, fail = createScriptFile(script)
	suite.Require().NoError(fail.ToError(), "should create file")

	content, fail := fileutils.ReadFile(suite.scriptFile.Filename())
	suite.Require().NoError(fail.ToError(), "unexpected error reading file contents")
	suite.Equal(script.Value(), string(content))
}

func (suite *EditTestSuite) TestGetOpenCmd_EditorSet() {
	expected := "debug"
	if runtime.GOOS == "windows" {
		expected = "debug.exe"
	}

	f, err := os.OpenFile(expected, os.O_CREATE|os.O_EXCL, 0700)
	suite.NoError(err, "should be able to create executable file")
	defer os.Remove(f.Name())

	err = f.Close()
	suite.NoError(err, "could no close file")

	originalPath := os.Getenv("PATH")
	defer os.Setenv("PATH", originalPath)

	wd, err := os.Getwd()
	suite.NoError(err, "could not get current working directory")

	err = os.Setenv("PATH", wd)
	suite.NoError(err, "could not set PATH")

	os.Setenv("EDITOR", expected)

	actual, fail := getOpenCmd()
	suite.Require().NoError(fail.ToError(), "could not get open command")
	suite.Equal(expected, actual)
}

func (suite *EditTestSuite) TestGetOpenCmd_EditorSet_NotInPath() {
	os.Setenv("EDITOR", "NotInPath")

	_, fail := getOpenCmd()
	suite.Require().Error(fail, "should get failure when editor is not in PATH")
}

func (suite *EditTestSuite) TestGetOpenCmd_EditorSet_InvalidFilepath() {
	wd, err := os.Getwd()
	suite.NoError(err, "could not get current working directory")

	executeable := "someExecutable"
	if runtime.GOOS == "windows" {
		executeable = "someExecutable.exe"
	}
	os.Setenv("EDITOR", filepath.Join(wd, executeable))

	_, fail := getOpenCmd()
	suite.Require().Error(fail, "should get failure when editor in path does not exist")
}

func (suite *EditTestSuite) TestGetOpenCmd_EditorSet_NoExtensionWindows() {
	if runtime.GOOS != "windows" {
		suite.T().Skip("the test for file extensions is only relevant for Windows")
	}

	wd, err := os.Getwd()
	suite.NoError(err, "could not get current working director")

	os.Setenv("EDITOR", filepath.Join(wd, "executable"))

	_, fail := getOpenCmd()
	suite.Require().Error(fail, "should get failure when editor path does not have extension")
}

func (suite *EditTestSuite) TestGetOpenCmd_EditorNotSet() {
	os.Setenv("EDITOR", "")
	var expected string
	platform := runtime.GOOS
	switch platform {
	case "linux":
		expected = openCmdLin
	case "darwin":
		expected = openCmdMac
	case "windows":
		expected = defaultEditorWin
	}

	actual, fail := getOpenCmd()
	if platform == "linux" && fail != nil {
		suite.EqualError(fail.ToError(), locale.Tr("error_open_not_installed_lin", openCmdLin))
	} else {
		suite.Require().NoError(fail.ToError(), "could not get open command")
		suite.Equal(expected, actual)
	}
}

func (suite *EditTestSuite) TestNewScriptWatcher() {
	script := suite.project.ScriptByName("hello")

	var fail *failures.Failure
	suite.scriptFile, fail = createScriptFile(script)
	suite.Require().NoError(fail.ToError(), "should create file")

	watcher, fail := newScriptWatcher(suite.scriptFile)
	suite.Require().NoError(fail.ToError(), "unexpected error creatig script watcher")

	go watcher.run()

	watcher.done <- true

	select {
	case fail = <-watcher.fails:
		suite.Require().NoError(fail.ToError(), "should not get error from running watcher")
	default:
		// Do nothing, test passed
	}
}

func (suite *EditTestSuite) TestUpdateProjectFile() {
	replace := suite.project.ScriptByName("replace")

	var fail *failures.Failure
	suite.scriptFile, fail = createScriptFile(replace)
	suite.Require().NoError(fail.ToError(), "unexpected error creating script file")

	EditArgs.Name = "hello"
	fail = updateProjectFile(suite.scriptFile)
	suite.Require().NoError(fail.ToError(), "should be able to update script file")

	updatedProject := project.Get()
	suite.Equal(replace.Value(), updatedProject.ScriptByName("hello").Value())
}

func TestEditSuite(t *testing.T) {
	suite.Run(t, new(EditTestSuite))
}
