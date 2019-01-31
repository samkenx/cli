package project

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/ActiveState/cli/internal/constraints"

	"github.com/ActiveState/cli/internal/failures"
	"github.com/ActiveState/cli/internal/variables"
	"github.com/ActiveState/cli/pkg/projectfile"
)

// FailProjectNotLoaded identifies a failure as being due to a missing project file
var FailProjectNotLoaded = failures.Type("project.fail.notparsed", failures.FailUser)

// Build covers the build structure
type Build map[string]string

// Project covers the platform structure
type Project struct {
	projectfile *projectfile.Project
}

// Platform covers the platform structure
type Platform struct {
	platform    *projectfile.Platform
	projectfile *projectfile.Project
}

// Language covers the language structure
type Language struct {
	language    *projectfile.Language
	projectfile *projectfile.Project
}

// Package covers the package structure
type Package struct {
	pkg         *projectfile.Package
	projectfile *projectfile.Project
}

// Script covers the script structure
type Script struct {
	script      *projectfile.Script
	projectfile *projectfile.Project
}

// Event covers the event structure
type Event struct {
	event       *projectfile.Event
	projectfile *projectfile.Project
}

// Variable covers the variable structure
type Variable struct {
	variable    *projectfile.Variable
	projectfile *projectfile.Project
}

// Build covers the build structure
type Build map[string]string

// Source returns the source projectfile
func (p *Project) Source() *projectfile.Project { return p.projectfile }

// Platforms gets platforms
func (p *Project) Platforms() []*Platform {
	platforms := []*Platform{}
	for i := range p.projectfile.Platforms {
		platforms = append(platforms, &Platform{&p.projectfile.Platforms[i], p.projectfile})
	}
	return platforms
}

// Languages returns a reference to projectfile.Languages
func (p *Project) Languages() []*Language {
	languages := []*Language{}
	for i, language := range p.projectfile.Languages {
		if !constraints.IsConstrained(language.Constraints) {
			languages = append(languages, &Language{&p.projectfile.Languages[i], p.projectfile})
		}
	}
	return languages
}

// Variables returns a reference to projectfile.Variables
func (p *Project) Variables() []*Variable {
	variables := []*Variable{}
	for i, variable := range p.projectfile.Variables {
		if !constraints.IsConstrained(variable.Constraints) {
			variables = append(variables, &Variable{&p.projectfile.Variables[i], p.projectfile})
		}
	}
	return variables
}

// Events returns a reference to projectfile.Events
func (p *Project) Events() []*Event {
	events := []*Event{}
	for i, event := range p.projectfile.Events {
		if !constraints.IsConstrained(event.Constraints) {
			events = append(events, &Event{&p.projectfile.Events[i], p.projectfile})
		}
	}
	return events
}

// Scripts returns a reference to projectfile.Scripts
func (p *Project) Scripts() []*Script {
	scripts := []*Script{}
	for i, script := range p.projectfile.Scripts {
		if !constraints.IsConstrained(script.Constraints) {
			scripts = append(scripts, &Script{&p.projectfile.Scripts[i], p.projectfile})
		}
	}
	return scripts
}

// Name returns project name
func (p *Project) Name() string { return p.projectfile.Name }

// NormalizedName returns the project name in a normalized format (alphanumeric, lowercase)
func (p *Project) NormalizedName() string {
	rx, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		failures.Handle(err, fmt.Sprintf("Regex failed to compile, error: %v", err))

		// This should only happen while in development, hence the os.Exit
		os.Exit(1)
	}

	return strings.ToLower(rx.ReplaceAllString(p.Name(), ""))
}

// Owner returns project owner
func (p *Project) Owner() string { return p.projectfile.Owner }

// Version returns project version
func (p *Project) Version() string { return p.projectfile.Version }

// Namespace returns project namespace
func (p *Project) Namespace() string { return p.projectfile.Namespace }

// Environments returns project environment
func (p *Project) Environments() string { return p.projectfile.Environments }

// Get returns project struct. Quits execution if error occurs
func Get() *Project {
	pj := projectfile.Get()
	return &Project{pj}
}

// GetSafe returns project struct.  Produces failure if error occurs, allows recovery
func GetSafe() (*Project, *failures.Failure) {
	pj, fail := projectfile.GetSafe()
	if fail.ToError() != nil {
		return nil, fail
	}
	return &Project{pj}, nil
}

// Platform covers the platform structure
type Platform struct {
	platform    *projectfile.Platform
	projectfile *projectfile.Project
}

// Source returns the source projectfile
func (p *Platform) Source() *projectfile.Project { return p.projectfile }

// Name returns platform name
func (p *Platform) Name() string { return p.platform.Name }

// Os returned with all variables evaluated
func (p *Platform) Os() string {
	value := variables.ExpandFromProject(p.platform.Os, p.projectfile)
	return value
}

// Version returned with all variables evaluated
func (p *Platform) Version() string {
	value := variables.ExpandFromProject(p.platform.Version, p.projectfile)
	return value
}

// Architecture with all variables evaluated
func (p *Platform) Architecture() string {
	value := variables.ExpandFromProject(p.platform.Architecture, p.projectfile)
	return value
}

// Libc returned are constrained and all variables evaluated
func (p *Platform) Libc() string {
	value := variables.ExpandFromProject(p.platform.Libc, p.projectfile)
	return value
}

// Compiler returned are constrained and all variables evaluated
func (p *Platform) Compiler() string {
	value := variables.ExpandFromProject(p.platform.Compiler, p.projectfile)
	return value
}

// Language covers the language structure
type Language struct {
	language    *projectfile.Language
	projectfile *projectfile.Project
}

// Source returns the source projectfile
func (l *Language) Source() *projectfile.Project { return l.projectfile }

// Name with all variables evaluated
func (l *Language) Name() string { return l.language.Name }

// Version with all variables evaluated
func (l *Language) Version() string { return l.language.Version }

// Build with all variables evaluated
func (l *Language) Build() *Build {
	build := Build{}
	for key, val := range l.language.Build {
		newVal := variables.ExpandFromProject(val, l.projectfile)
		build[key] = newVal
	}
	return &build
}

// Packages returned are constrained set
func (l *Language) Packages() []Package {
	validPackages := []Package{}
	for i, pkg := range l.language.Packages {
		if !constraints.IsConstrained(pkg.Constraints) {
			newPkg := Package{}
			newPkg.pkg = &l.language.Packages[i]
			newPkg.projectfile = l.projectfile
			validPackages = append(validPackages, newPkg)
		}
	}
	return validPackages
}

// Package covers the package structure
type Package struct {
	pkg         *projectfile.Package
	projectfile *projectfile.Project
}

// Source returns the source projectfile
func (p *Package) Source() *projectfile.Project { return p.projectfile }

// Name returns package name
func (p *Package) Name() string { return p.pkg.Name }

// Version returns package version
func (p *Package) Version() string { return p.pkg.Version }

// Build returned with all variables evaluated
func (p *Package) Build() *Build {
	build := Build{}
	for key, val := range p.pkg.Build {
		newVal := variables.ExpandFromProject(val, p.projectfile)
		build[key] = newVal
	}
	return &build
}

// Variable covers the variable structure
type Variable struct {
	variable    *projectfile.Variable
	projectfile *projectfile.Project
}

// Source returns the source projectfile
func (v *Variable) Source() *projectfile.Project { return v.projectfile }

// Name returns variable name
func (v *Variable) Name() string { return v.variable.Name }

// Value returned with all variables evaluated
func (v *Variable) Value() string {
	if v.variable.Value != nil {
		value := variables.ExpandFromProject(*v.variable.Value, v.projectfile)
		return value
	} else {
		return ""
	}
}

// Value returned with all variables evaluated
func (v *Variable) Location() string {
	if v.variable.Location != nil {
		value := variables.ExpandFromProject(*v.variable.Location, v.projectfile)
		return value
	} else {
		return ""
	}
}

// Hook covers the hook structure
type Hook struct {
	hook        *projectfile.Hook
	projectfile *projectfile.Project
}

// Source returns the source projectfile
func (e *Event) Source() *projectfile.Project { return e.projectfile }

// Name returns Event name
func (e *Event) Name() string { return e.event.Name }

// Value returned with all variables evaluated
func (e *Event) Value() string {
	value := variables.ExpandFromProject(e.event.Value, e.projectfile)
	return value
}

// Script covers the command structure
type Script struct {
	script      *projectfile.Script
	projectfile *projectfile.Project
}

// Source returns the source projectfile
func (script *Script) Source() *projectfile.Project { return script.projectfile }

// Name returns script name
func (script *Script) Name() string { return script.script.Name }

// Value returned with all variables evaluated
func (script *Script) Value() string {
	value := variables.ExpandFromProject(script.script.Value, script.projectfile)
	return value
}

// Standalone returns if the script is standalone or not
func (script *Script) Standalone() bool { return script.script.Standalone }
