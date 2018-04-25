package projectfile

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/ActiveState/cli/internal/constants"
	"github.com/ActiveState/cli/internal/failures"
	"github.com/ActiveState/cli/internal/locale"
	"github.com/ActiveState/cli/internal/logging"
	"github.com/hashicorp/hcl"
	"github.com/hashicorp/hcl/hcl/ast"
	"github.com/hashicorp/hcl/hcl/printer"
	"github.com/mitchellh/hashstructure"
)

// FailNoProject identifies a failure as being due to a missing project file
var FailNoProject = failures.Type("projectfile.fail.noproject")

// Project covers the top level project structure of our hcl
type Project struct {
	Name         string              `hcl:"name"`
	Owner        string              `hcl:"owner"`
	Namespace    string              `hcl:"namespace"`
	Version      string              `hcl:"version"`
	Environments string              `hcl:"environments"`
	Platforms    map[string]Platform `hcl:"platform"`
	Languages    map[string]Language `hcl:"language"`
	Variables    map[string]Variable `hcl:"variable"`
	Hooks        []Hook              `hcl:"hook"`
	Commands     map[string]Command  `hcl:"command"`
	path         string              // "private"
	astFile      *ast.File           // "private"
}

// Platform covers the platform structure of our hcl
type Platform struct {
	Os           string `hcl:"os"`
	Version      string `hcl:"version"`
	Architecture string `hcl:"architecture"`
	Libc         string `hcl:"libc"`
	Compiler     string `hcl:"compiler"`
}

// Build covers the build map, which can go under languages or packages
// Build can hold variable keys, so we cannot predict what they are, hence why it is a map
type Build map[string]string

// Language covers the language structure, which goes under Project
type Language struct {
	Version     string             `hcl:"version"`
	Constraints Constraint         `hcl:"constraint"`
	Build       Build              `hcl:"build"`
	Packages    map[string]Package `hcl:"package"`
}

// Constraint covers the constraint structure, which can go under almost any other struct
type Constraint struct {
	Platform    string `hcl:"platform"`
	Environment string `hcl:"environment"`
}

// Package covers the package structure, which goes under the language struct
type Package struct {
	Version     string     `hcl:"version"`
	Constraints Constraint `hcl:"constraint"`
	Build       Build      `hcl:"build"`
}

// Variable covers the variable structure, which goes under Project
type Variable struct {
	Value       string     `hcl:"value"`
	Constraints Constraint `hcl:"constraint"`
}

// Hook covers the hook structure, which goes under Project
type Hook struct {
	Name        string     `hcl:"name"`
	Value       string     `hcl:"value"`
	Constraints Constraint `hcl:"constraint"`
}

// Hash return a hashed version of the hook
func (h *Hook) Hash() (string, error) {
	hash, err := hashstructure.Hash(h, nil)
	if err != nil {
		logging.Errorf("Cannot hash hook: %v", err)
		return "", err
	}
	return fmt.Sprintf("%X", hash), nil
}

// Command covers the command structure, which goes under Project
type Command struct {
	Value       string     `hcl:"value"`
	Constraints Constraint `hcl:"constraint"`
}

var persistentProject *Project

// Parse the given filepath, which should be the full path to an activestate.hcl file
func Parse(filepath string) (*Project, error) {
	dat, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	project := Project{}
	project.astFile, err = hcl.ParseBytes(dat)
	project.path = filepath

	if err != nil {
		return nil, FailNoProject.New(locale.T("err_project_parse", map[string]interface{}{"Error": err.Error()}))
	}

	err = hcl.DecodeObject(&project, project.astFile)

	if err != nil {
		return nil, FailNoProject.New(locale.T("err_project_parse", map[string]interface{}{"Error": err.Error()}))
	}

	project.FixUnmarshalledHooks() // temporary workaround for HCL bug

	return &project, err
}

// Path returns the project's activestate.hcl file path.
func (p *Project) Path() string {
	return p.path
}

// SetPath sets the path of the project file and should generally only be used by tests
func (p *Project) SetPath(path string) {
	p.path = path
}

// Save the project to its activestate.hcl file
func (p *Project) Save() error {
	// TODO: update p.astNode nodes with project fields
	//dat, err := hcl.Marshal(p)
	//if err != nil {
	//	return err
	//}

	f, err := os.Create(p.Path())
	defer f.Close()
	if err != nil {
		return err
	}

	err = printer.Fprint(f, p.astFile)
	if err != nil {
		return err
	}

	return nil
}

// Returns the path to the project activestate.hcl
func getProjectFilePath() string {
	root, err := os.Getwd()
	if err != nil {
		logging.Warning("Could not get project root path: %v", err)
		return ""
	}
	return filepath.Join(root, constants.ConfigFileName)
}

// Get returns the project configration in an unsafe manner (exits if errors occur)
func Get() *Project {
	project, err := GetSafe()
	if err != nil {
		failures.Handle(err, locale.T("err_project_file_unavailable"))
		os.Exit(1)
	}

	return project
}

// GetSafe returns the project configuration in a safe manner (returns error)
func GetSafe() (*Project, error) {
	if persistentProject != nil {
		return persistentProject, nil
	}

	projectFilePath := os.Getenv(constants.ProjectEnvVarName)
	if projectFilePath == "" {
		projectFilePath = getProjectFilePath()
	}

	_, err := ioutil.ReadFile(projectFilePath)
	if err != nil {
		logging.Warning("Cannot load config file: %v", err)
		return nil, FailNoProject.New(locale.T("err_no_projectfile"))
	}
	project, err := Parse(projectFilePath)
	if err == nil {
		project.Persist()
	}
	return project, err
}

// Reset the current state, which unsets the persistent project
func Reset() {
	persistentProject = nil
	os.Unsetenv(constants.ProjectEnvVarName)
}

// Persist "activates" the given project and makes it such that subsequent calls
// to Get() return this project.
// Only one project can persist at a time.
func (p *Project) Persist() {
	persistentProject = p
	os.Setenv(constants.ProjectEnvVarName, p.Path())
}

// FixUnmarshalledHooks is a temporary workaround to fix HCL-decoded hook
// arrays. See: https://github.com/hashicorp/hcl/issues/247.
// This function needs to be called on any &project passed to `hcl.Unmarshal()`
// or hcl.Decode*() that contains multiple `hook { ... }` definitions, which
// usually only happens in test functions.
// This function is automatically called by functions in this package.
func (p *Project) FixUnmarshalledHooks() {
	fixedHooks := []Hook{}
	fixedHook := Hook{}
	for _, badHook := range p.Hooks {
		if badHook.Name != "" {
			fixedHook.Name = badHook.Name
		} else if badHook.Value != "" {
			fixedHook.Value = badHook.Value
		}
		if fixedHook.Name != "" && fixedHook.Value != "" {
			fixedHooks = append(fixedHooks, fixedHook)
			fixedHook = Hook{}
		}
	}
	p.Hooks = fixedHooks
}
