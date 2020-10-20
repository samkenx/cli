package secrets

import (
	"fmt"

	"github.com/ActiveState/cli/internal/access"
	"github.com/ActiveState/cli/internal/failures"
	"github.com/ActiveState/cli/internal/locale"
	"github.com/ActiveState/cli/internal/logging"
	"github.com/ActiveState/cli/internal/output"
	"github.com/ActiveState/cli/internal/primer"
	"github.com/ActiveState/cli/internal/secrets"
	secretsapi "github.com/ActiveState/cli/pkg/platform/api/secrets"
	secretsModels "github.com/ActiveState/cli/pkg/platform/api/secrets/secrets_models"
	"github.com/ActiveState/cli/pkg/project"
)

type listPrimeable interface {
	primer.Outputer
	primer.Projecter
}

// ListRunParams tracks the info required for running List.
type ListRunParams struct {
	Filter string
}

// List manages the listing execution context.
type List struct {
	secretsClient *secretsapi.Client
	out           output.Outputer
	proj          *project.Project
}

// NewList prepares a list execution context for use.
func NewList(client *secretsapi.Client, p listPrimeable) *List {
	return &List{
		secretsClient: client,
		out:           p.Output(),
		proj:          p.Project(),
	}
}

// Run executes the list behavior.
func (l *List) Run(params ListRunParams) error {
	if err := checkSecretsAccess(l.proj); err != nil {
		return locale.WrapError(err, "secrets_err")
	}

	defs, fail := definedSecrets(l.proj, l.secretsClient, params.Filter)
	if fail != nil {
		return locale.WrapError(fail, "secrets_err_defined")
	}

	exports, fail := defsToSecrets(defs)
	if fail != nil {
		return locale.WrapError(fail, "secrets_err_values")
	}

	l.out.Print(exports)

	return nil
}

func checkSecretsAccess(proj *project.Project) error {
	allowed, fail := access.Secrets(proj.Owner())
	if fail != nil {
		return locale.WrapError(fail, "secrets_err_access")
	}
	if !allowed {
		return locale.NewError("secrets_warning_no_access")
	}
	return nil
}

func definedSecrets(proj *project.Project, secCli *secretsapi.Client, filter string) ([]*secretsModels.SecretDefinition, *failures.Failure) {
	logging.Debug("listing variables for org=%s, project=%s", proj.Owner(), proj.Name())

	secretDefs, fail := secrets.DefsByProject(secCli, proj.Owner(), proj.Name())
	if fail != nil {
		return nil, fail
	}

	if filter != "" {
		secretDefs = filterSecrets(proj, secretDefs, filter)
	}

	return secretDefs, nil
}

func filterSecrets(proj *project.Project, secrectDefs []*secretsModels.SecretDefinition, filter string) []*secretsModels.SecretDefinition {
	secrectDefsFiltered := []*secretsModels.SecretDefinition{}

	oldExpander := project.RegisteredExpander("secrets")
	if oldExpander != nil {
		defer project.RegisterExpander("secrets", oldExpander)
	}
	expander := project.NewSecretExpander(secretsapi.Get(), proj)
	project.RegisterExpander("secrets", expander.Expand)
	project.ExpandFromProject(fmt.Sprintf("$%s", filter), proj)
	accessedSecrets := expander.SecretsAccessed()
	if accessedSecrets == nil {
		return secrectDefsFiltered
	}

	for _, secretDef := range secrectDefs {
		isUser := *secretDef.Scope == secretsModels.SecretDefinitionScopeUser
		for _, accessedSecret := range accessedSecrets {
			if accessedSecret.Name == *secretDef.Name && accessedSecret.IsUser == isUser {
				secrectDefsFiltered = append(secrectDefsFiltered, secretDef)
			}
		}
	}

	return secrectDefsFiltered
}

func defsToSecrets(defs []*secretsModels.SecretDefinition) ([]*SecretListExport, *failures.Failure) {
	secretsExport := make([]*SecretListExport, len(defs))
	expander := project.NewSecretExpander(secretsapi.Get(), project.Get())

	for i, def := range defs {
		if def.Name == nil || def.Scope == nil {
			logging.Error("Could not get pointer for secret name and/or scope, definition ID: %d", def.DefID)
			continue
		}

		secretValue, fail := expander.FindSecret(*def.Name, *def.Scope == secretsModels.SecretDefinitionScopeUser)
		if fail != nil {
			return secretsExport, fail
		}

		desc := "-"
		if def.Description != "" {
			desc = def.Description
		}

		defStatus := locale.T("secrets_row_value_unset")
		if secretValue != nil && secretValue.Value != nil {
			defStatus = locale.T("secrets_row_value_set")
		}

		secretsExport[i] = &SecretListExport{
			Name:        *def.Name,
			Scope:       *def.Scope,
			Value:       defStatus,
			Description: desc,
			Usage:       *def.Scope + "." + *def.Name,
		}
	}

	return secretsExport, nil
}

// SecretListExport defines important information about a secret that should be
// displayed.
type SecretListExport struct {
	Name        string `json:"name"`
	Scope       string `json:"scope"`
	Value       string `json:"value"` // AKA "Definition Status - Defined/Undefined"
	Description string `json:"description"`
	Usage       string `json:"usage"`
}

type xSecretGetExport struct {
	Name        string `json:"name"`
	Scope       string `json:"scope"`
	Description string `json:"description"`
	HasValue    bool   `json:"has_value"`
	Value       string `json:"value,omitempty"`
}
