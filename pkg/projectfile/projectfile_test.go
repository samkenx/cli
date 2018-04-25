package projectfile

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/ActiveState/cli/internal/constants"
	"github.com/ActiveState/cli/internal/environment"
	"github.com/hashicorp/hcl"
	"github.com/stretchr/testify/assert"
)

func TestProjectStruct(t *testing.T) {
	project := Project{}
	dat := strings.TrimSpace(`
		name = "valueForName"
		owner = "valueForOwner"
		namespace = "valueForNamespace"
		version = "valueForVersion"
		environments = "valueForEnvironments"
	`)

	err := hcl.Unmarshal([]byte(dat), &project)
	assert.Nil(t, err, "Should not throw an error")

	assert.Equal(t, "valueForName", project.Name, "Name should be set")
	assert.Equal(t, "valueForOwner", project.Owner, "Owner should be set")
	assert.Equal(t, "valueForNamespace", project.Namespace, "Namespace should be set")
	assert.Equal(t, "valueForVersion", project.Version, "Version should be set")
	assert.Equal(t, "valueForEnvironments", project.Environments, "Environments should be set")
	assert.Equal(t, "", project.Path(), "Path should be empty")
}

func TestPlatformStruct(t *testing.T) {
	root := []map[string]map[string]Platform{}
	dat := strings.TrimSpace(`
		platform "valueForName" {
			os = "valueForOS"
			version = "valueForVersion"
			architecture = "valueForArch"
		}
	`)

	err := hcl.Unmarshal([]byte(dat), &root)
	assert.Nil(t, err, "Should not throw an error")
	assert.Equal(t, 1, len(root), "One config node was detected")

	platforms := root[0]["platform"]
	assert.NotNil(t, platforms, "At least one platform was detected")
	assert.Equal(t, 1, len(platforms), "One platform was defined")

	assert.NotNil(t, platforms["valueForName"], "Name should be set")
	assert.Equal(t, "valueForOS", platforms["valueForName"].Os, "OS should be set")
	assert.Equal(t, "valueForVersion", platforms["valueForName"].Version, "Version should be set")
	assert.Equal(t, "valueForArch", platforms["valueForName"].Architecture, "Architecture should be set")
}

func TestBuildStruct(t *testing.T) {
	root := []map[string]Build{}
	dat := strings.TrimSpace(`
		build {
			key1 = "val1"
			key2 = "val2"
		}
	`)

	err := hcl.Unmarshal([]byte(dat), &root)
	assert.Nil(t, err, "Should not throw an error")
	assert.Equal(t, 1, len(root), "One config node was detected")

	build := root[0]["build"]
	assert.NotNil(t, build, "Build options were detected")

	assert.Equal(t, "val1", build["key1"], "Key1 should be set")
	assert.Equal(t, "val2", build["key2"], "Key2 should be set")
}

func TestLanguageStruct(t *testing.T) {
	root := []map[string]map[string]Language{}
	dat := strings.TrimSpace(`
		language "valueForName" {
			version = "valueForVersion"
		}
	`)

	err := hcl.Unmarshal([]byte(dat), &root)
	assert.Nil(t, err, "Should not throw an error")
	assert.Equal(t, 1, len(root), "One config node was detected")

	languages := root[0]["language"]
	assert.NotNil(t, languages, "At least one language was detected")
	assert.Equal(t, 1, len(languages), "One language was defined")

	assert.NotNil(t, languages["valueForName"], "Name should be set")
	assert.Equal(t, "valueForVersion", languages["valueForName"].Version, "Version should be set")
}

func TestConstraintStruct(t *testing.T) {
	root := []map[string]Constraint{}
	dat := strings.TrimSpace(`
		constraint {
			platform = "valueForPlatform"
			environment = "valueForEnvironment"
		}
	`)

	err := hcl.Unmarshal([]byte(dat), &root)
	assert.Nil(t, err, "Should not throw an error")
	assert.Equal(t, 1, len(root), "One config node was detected")

	constraint := root[0]["constraint"]
	assert.NotNil(t, constraint, "A constraint was detected")

	assert.Equal(t, "valueForPlatform", constraint.Platform, "Platform should be set")
	assert.Equal(t, "valueForEnvironment", constraint.Environment, "Environment should be set")
}

func TestPackageStruct(t *testing.T) {
	root := []map[string]map[string]Package{}
	dat := strings.TrimSpace(`
		package "valueForName" {
			version = "valueForVersion"
		}
	`)

	err := hcl.Unmarshal([]byte(dat), &root)
	assert.Nil(t, err, "Should not throw an error")
	assert.Equal(t, 1, len(root), "One config node was detected")

	packages := root[0]["package"]
	assert.NotNil(t, packages, "At least one package was detected")
	assert.Equal(t, 1, len(packages), "One package was defined")

	assert.NotNil(t, packages["valueForName"], "Name should be set")
	assert.Equal(t, "valueForVersion", packages["valueForName"].Version, "Version should be set")
}

func TestVariableStruct(t *testing.T) {
	root := []map[string]map[string]Variable{}
	dat := strings.TrimSpace(`
		variable "valueForName" {
			value = "valueForValue"
		}
	`)

	err := hcl.Unmarshal([]byte(dat), &root)
	assert.Nil(t, err, "Should not throw an error")
	assert.Equal(t, 1, len(root), "One config node was detected")

	variables := root[0]["variable"]
	assert.NotNil(t, variables, "At least one variable was detected")
	assert.Equal(t, 1, len(variables), "One variable was defined")

	assert.NotNil(t, variables["valueForName"], "Name should be set")
	assert.Equal(t, "valueForValue", variables["valueForName"].Value, "Value should be set")
}

func TestHookStruct(t *testing.T) {
	root := []map[string]Hook{}
	dat := strings.TrimSpace(`
		hook {
			name = "valueForName"
			value = "valueForValue"
		}

		hook {
			name = "valueForName"
			value = "valueForValue2"
		}
	`)

	err := hcl.Unmarshal([]byte(dat), &root)
	assert.Nil(t, err, "Should not throw an error")
	assert.Equal(t, 2, len(root), "Two config nodes were detected")

	hook := root[0]["hook"]
	assert.NotNil(t, hook, "A hook was detected")

	assert.Equal(t, "valueForName", hook.Name, "Name should be set")
	assert.Equal(t, "valueForValue", hook.Value, "Value should be set")

	hook = root[1]["hook"]
	assert.NotNil(t, hook, "Another hook was detected")
	assert.Equal(t, "valueForName", hook.Name, "Name should be set")
	assert.Equal(t, "valueForValue2", hook.Value, "Value should be set")
}

func TestCommandStruct(t *testing.T) {
	root := []map[string]map[string]Command{}
	dat := strings.TrimSpace(`
		command "valueForName" {
			value = "valueForCommand"
		}
	`)

	err := hcl.Unmarshal([]byte(dat), &root)
	assert.Nil(t, err, "Should not throw an error")
	assert.Equal(t, 1, len(root), "One config node was detected")

	commands := root[0]["command"]
	assert.NotNil(t, commands, "At least one command was detected")
	assert.Equal(t, 1, len(commands), "One command was defined")

	assert.NotNil(t, commands["valueForName"], "Name should be set")
	assert.Equal(t, "valueForCommand", commands["valueForName"].Value, "Command should be set")
}

func TestParse(t *testing.T) {
	rootpath, err := environment.GetRootPath()

	if err != nil {
		t.Fatal(err)
	}

	project, err := Parse(filepath.Join(rootpath, "activestate.hcl.nope"))
	fmt.Println(err)
	assert.NotNil(t, err, "Should throw an error")

	project, err = Parse(filepath.Join(rootpath, "test", "activestate.hcl"))
	assert.Nil(t, err, "Should not throw an error")

	assert.NotEmpty(t, project.Name, "Name should be set")
	assert.NotEmpty(t, project.Owner, "Owner should be set")
	assert.NotEmpty(t, project.Namespace, "Namespace should be set")
	assert.NotEmpty(t, project.Version, "Version should be set")
	assert.NotEmpty(t, project.Platforms, "Platforms should be set")
	assert.NotEmpty(t, project.Environments, "Environments should be set")

	assert.NotNil(t, project.Platforms["Linux64Label"], "Platform name should be set")
	assert.NotEmpty(t, project.Platforms["Linux64Label"].Os, "Platform OS name should be set")
	assert.NotEmpty(t, project.Platforms["Linux64Label"].Architecture, "Platform architecture name should be set")
	assert.NotEmpty(t, project.Platforms["Linux64Label"].Libc, "Platform libc name should be set")
	assert.NotEmpty(t, project.Platforms["Linux64Label"].Compiler, "Platform compiler name should be set")

	assert.NotNil(t, project.Languages["Go"], "Language name should be set")
	assert.NotEmpty(t, project.Languages["Go"].Version, "Language version should be set")

	assert.NotNil(t, project.Languages["Go"].Packages["golang.org/x/crypto"], "Package name should be set")
	assert.NotEmpty(t, project.Languages["Go"].Packages["golang.org/x/crypto"].Version, "Package version should be set")

	assert.NotEmpty(t, project.Languages["Go"].Packages["golang.org/x/crypto"].Build, "Package build should be set")
	assert.NotEmpty(t, project.Languages["Go"].Packages["golang.org/x/crypto"].Build["debug"], "Build debug should be set")

	assert.NotEmpty(t, project.Languages["Go"].Packages["gopkg.in/yaml.v2"].Build, "Package build should be set")
	assert.NotEmpty(t, project.Languages["Go"].Packages["gopkg.in/yaml.v2"].Build["override"], "Build override should be set")

	assert.NotEmpty(t, project.Languages["Go"].Constraints.Platform, "Platform constraint should be set")
	assert.NotEmpty(t, project.Languages["Go"].Constraints.Environment, "Environment constraint should be set")

	assert.NotNil(t, project.Variables["DEBUG"], "Variable name should be set")
	assert.NotEmpty(t, project.Variables["DEBUG"].Value, "Variable value should be set")

	assert.NotEmpty(t, project.Hooks[0].Name, "Hook name should be set")
	assert.NotEmpty(t, project.Hooks[0].Value, "Hook value should be set")
	assert.NotEmpty(t, project.Hooks[1].Name, "Hook name should be set")
	assert.NotEmpty(t, project.Hooks[1].Value, "Hook value should be set")

	assert.NotEmpty(t, project.Commands["tests"], "Command name should be set")
	assert.NotEmpty(t, project.Commands["tests"].Value, "Command value should be set")

	assert.NotEmpty(t, project.Path(), "Path should be set")
}

func TestSave(t *testing.T) {
	rootpath, err := environment.GetRootPath()

	if err != nil {
		t.Fatal(err)
	}

	path := filepath.Join(rootpath, "test", "activestate.hcl")
	project, err := Parse(path)
	assert.NoError(t, err, "Should parse our hcl file")

	tmpfile, err := ioutil.TempFile("", "test")
	assert.NoError(t, err, "Should create a temp file")

	project.path = tmpfile.Name()
	project.Save()

	stat, err := tmpfile.Stat()
	assert.NoError(t, err, "Should be able to stat file")

	err = tmpfile.Close()
	assert.NoError(t, err, "Should close our temp file")

	assert.FileExists(t, tmpfile.Name(), "Project file is saved")
	assert.NotZero(t, stat.Size(), "Project file should have data")

	os.Remove(tmpfile.Name())
}

// Call getProjectFilePath
func TestGetProjectFilePath(t *testing.T) {
	root, err := environment.GetRootPath()
	assert.NoError(t, err, "Should detect root path")
	cwd, err := os.Getwd()
	assert.NoError(t, err, "Should fetch cwd")
	os.Chdir(filepath.Join(root, "test"))

	configPath := getProjectFilePath()
	expectedPath := filepath.Join(root, "test", constants.ConfigFileName)
	assert.Equal(t, expectedPath, configPath, "Project path is properly detected")

	os.Chdir(cwd) // restore
}

// Call getProjectFilePath but doesn't exist
func TestGetFail(t *testing.T) {
	config, _ := GetSafe()
	assert.Nil(t, config, "Config should not be set.")
	assert.Equal(t, "", os.Getenv(constants.ProjectEnvVarName), "The state should not be activated")

	Reset()
}

// TestGet the config
func TestGet(t *testing.T) {
	root, err := environment.GetRootPath()
	assert.NoError(t, err, "Should detect root path")
	cwd, _ := os.Getwd()
	os.Chdir(filepath.Join(root, "test"))

	config := Get()
	assert.NotNil(t, config, "Config should be set")
	assert.NotEqual(t, "", os.Getenv(constants.ProjectEnvVarName), "The project env var should be set")

	os.Chdir(cwd) // restore

	Reset()
}

func TestGetActivated(t *testing.T) {
	root, _ := environment.GetRootPath()
	cwd, _ := os.Getwd()
	os.Chdir(filepath.Join(root, "test"))

	config1 := Get()
	assert.Equal(t, filepath.Join(root, "test", constants.ConfigFileName), os.Getenv(constants.ProjectEnvVarName), "The activated state's config file is set")

	os.Chdir(root)
	config2, err := GetSafe()
	assert.NoError(t, err, "No error even if no activestate.hcl does not exist")
	assert.Equal(t, config1, config2, "The same activated state is returned")

	expected := filepath.Join(root, "test", constants.ConfigFileName)
	actual := os.Getenv(constants.ProjectEnvVarName)
	assert.Equal(t, expected, actual, "The activated state's config file is still set properly")

	os.Chdir(cwd) // restore

	Reset()
}
