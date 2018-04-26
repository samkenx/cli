package projectfile

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

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
// This is done by updating AST nodes for values that have changed and adding
// AST nodes for new values.
func (p *Project) Save() error {
	// Update root AST nodes.
	for key, value := range map[string]string{
		"name":         p.Name,
		"owner":        p.Owner,
		"namespace":    p.Namespace,
		"version":      p.Version,
		"environments": p.Environments,
	} {
		p.updateASTRootNodeValue(key, value)
	}

	// Update AST platform nodes.
	//p.updateASTPlatformNodes()

	// Update AST language nodes.
	//p.updateASTLanguageNodes()

	// Update AST variable nodes.
	//p.updateASTVariableNodes()

	// Update AST hook nodes.
	p.updateASTHookNodes()

	// Update AST command nodes.
	//p.updateASTCommandNodes()

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

// Updates the string value of a root node in the HCL config file AST.
func (p *Project) updateASTRootNodeValue(key, value string) {
	keyList := p.astFile.Node.(*ast.ObjectList).Filter(key)
	if keyList.Items == nil {
		logging.Debug("Could not update AST root node value for '%s': key not found", key)
		return // benign programming error
	}
	if nodeValue, ok := keyList.Items[0].Val.(*ast.LiteralType); ok && strings.Trim(nodeValue.Token.Text, "\"") != value {
		nodeValue.Token.Text = fmt.Sprintf("\"%s\"", value)
	}
}

func (p *Project) updateASTHookNodes() {
	logging.Debug("Updating AST hook nodes")

	// Track project hooks that need to be added to the HCL config AST.
	// Start with the premise that all project hooks need to be added. The list
	// will be trimmed as equivalent AST nodes are found.
	added := make([]Hook, len(p.Hooks))
	copy(added, p.Hooks)

	// Track HCL config AST nodes that need to be removed.
	removed := []*ast.ObjectType{}

	// Loop through all AST hook nodes.
	// Any invalid nodes are skipped (not removed) so-as to not be destructive
	// with a user's config file.
	hookList := p.astFile.Node.(*ast.ObjectList).Filter("hook")
	if hookList.Items == nil {
		hookList.Items = []*ast.ObjectItem{}
	}
	for i, item := range hookList.Items {
		// Ensure the hook node to check is an object.
		node, ok := item.Val.(*ast.ObjectType)
		if !ok {
			logging.Debug("Hook node #%d is invalid (not an object)", i)
			continue // skip invalid hook
		}
		logging.Debug("Attempting to find a project hook for hook node #%d", i)

		// Track project hooks that may match with this hook node.
		potentialMatches := make([]Hook, len(p.Hooks))
		copy(potentialMatches, p.Hooks)

		// Fetch the hook node's keys and filter out potential hook matches with
		// different values.
		nextNode := false
		for _, key := range []string{"name", "value"} {
			keyList := node.List.Filter(key)
			if keyList.Items == nil {
				logging.Debug("Hook node #%d is invalid (no %s key)", i, key)
				nextNode = true
				break
			}
			nodeValue, ok := keyList.Items[0].Val.(*ast.LiteralType)
			if !ok {
				logging.Debug("Hook node #%d is invalid (%s not a string)", i, key)
				nextNode = true
				break
			}
			text := strings.Trim(nodeValue.Token.Text, "\"")
			for j := 0; j < len(potentialMatches); j++ {
				hook := potentialMatches[j]
				valueMap := map[string]string{
					"name":  hook.Name,
					"value": hook.Value,
				}
				if text != valueMap[key] {
					logging.Debug("Removing candidate project hook %v for hook node #%d (%s != '%s')", hook, i, key, text)
					potentialMatches = append(potentialMatches[:j], potentialMatches[j+1:]...)
					j-- // stay on the same element index
				}
			}
			// If no equivalent project hooks were found, this hook must have been
			// removed.
			if len(potentialMatches) == 0 {
				logging.Debug("Removing hook node #%d (no matching project hooks found)", i)
				removed = append(removed, node)
				nextNode = true
				break
			}
		}
		if nextNode {
			continue // either skip invalid hook or move to next node
		}

		// Fetch the hook node's "constraint" object and filter out potential hook
		// matches with different constraints.
		constraintList := node.List.Filter("constraint")
		for j := 0; j < len(potentialMatches); j++ {
			hook := potentialMatches[j]
			if constraintList.Items == nil {
				if hook.Constraints.Platform != "" || hook.Constraints.Environment != "" {
					logging.Debug("Removing candidate project hook %v for hook node #%d (non-empty constraint)", hook, i)
					potentialMatches = append(potentialMatches[:j], potentialMatches[j+1:]...)
					j-- // stay on the same element index
				}
			} else {
				constraintNode, ok := constraintList.Items[0].Val.(*ast.ObjectType)
				if !ok {
					logging.Debug("Hook node #%d is invalid (constraint not an object)", i)
					nextNode = true
					break
				}
				for name, value := range map[string]string{
					"platform":    hook.Constraints.Platform,
					"environment": hook.Constraints.Environment,
				} {
					nameList := constraintNode.List.Filter(name)
					if nameList.Items == nil && value != "" {
						logging.Debug("Removing candidate project hook %v for hook node #%d (constraint %s != '')", hook, i, name)
						potentialMatches = append(potentialMatches[:j], potentialMatches[j+1:]...)
						j--   // stay on the same element index
						break // will continue looping through potentialMatches
					}
					nodeValue, ok := nameList.Items[0].Val.(*ast.LiteralType)
					if !ok && value != "" {
						logging.Debug("Removing candidate project hook %v for hook node #%d (constraint %s != '')", hook, i, name)
						potentialMatches = append(potentialMatches[:j], potentialMatches[j+1:]...)
						j--   // stay on the same element index
						break // will continue looping through potentialMatches
					}
					text := strings.Trim(nodeValue.Token.Text, "\"")
					if text != value {
						logging.Debug("Removing candidate project hook %v for hook node #%d (constraint %s != '%s')", hook, i, name, text)
						potentialMatches = append(potentialMatches[:j], potentialMatches[j+1:]...)
						j--   // stay on the same element index
						break // will continue looping through potentialMatches
					}
				}
			}
		}
		if nextNode {
			continue
		}
		// If no project hooks were found, this hook must have been removed.
		if len(potentialMatches) == 0 {
			logging.Debug("Removing hook node #%d (no matching project hooks found)", i)
			removed = append(removed, node)
			continue
		}

		// At this point, a matching project hook was found for this hook node.
		// It does not need to be added as a new node.
		logging.Debug("Found matching project hook for hook node #%d: %v", i, potentialMatches[0])
		if len(potentialMatches) > 1 {
			logging.Debug("More than one project hook found for hook node #%d:", i)
			logging.Debug("%v", potentialMatches[1:])
		}
		for j, hook := range added {
			if hook == potentialMatches[0] {
				added = append(added[:j], added[j+1:]...)
			}
		}
	}

	nodes := p.astFile.Node.(*ast.ObjectList)

	// Remove hook nodes that were marked for removal.
	for _, node := range removed {
		for i, item := range nodes.Items {
			existingNode, ok := item.Val.(*ast.ObjectType)
			if !ok {
				continue
			}
			if existingNode == node {
				nodes.Items = append(nodes.Items[:i], nodes.Items[i+1:]...)
				break
			}
		}
	}

	// Add hook nodes that were marked for adding.
	for _, hook := range added {
		logging.Debug("Adding new hook node for project hook %v", hook)
		if hook.Constraints.Platform != "" || hook.Constraints.Environment != "" {
			node, _ := hcl.Parse(fmt.Sprintf(`
				hook {
					name = "%s"
					value = "%s"
					constraint {
						platform = "%s"
						environment = "%s"
					}
				}
			`, hook.Name, hook.Value, hook.Constraints.Platform, hook.Constraints.Environment))
			nodes.Add(node.Node.(*ast.ObjectList).Items[0])
		} else {
			node, _ := hcl.Parse(fmt.Sprintf(`
				hook {
					name = "%s"
					value = "%s"
				}
			`, hook.Name, hook.Value))
			nodes.Add(node.Node.(*ast.ObjectList).Items[0])
		}
	}
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
