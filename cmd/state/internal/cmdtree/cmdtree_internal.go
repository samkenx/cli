// +build !external

package cmdtree

import (
	"github.com/ActiveState/cli/internal/captain"
	"github.com/ActiveState/cli/internal/logging"
	secretsapi "github.com/ActiveState/cli/pkg/platform/api/secrets"
	"github.com/ActiveState/cli/state/events"
	"github.com/ActiveState/cli/state/export"
	"github.com/ActiveState/cli/state/fork"
	"github.com/ActiveState/cli/state/invite"
	pkg "github.com/ActiveState/cli/state/package"
	"github.com/ActiveState/cli/state/projects"
	"github.com/ActiveState/cli/state/pull"
	"github.com/ActiveState/cli/state/run"
	"github.com/ActiveState/cli/state/scripts"
	"github.com/ActiveState/cli/state/secrets"
	"github.com/ActiveState/cli/state/show"
	"github.com/ActiveState/cli/state/update"
)

// applyLegacyChildren will register any commands and expanders
func applyLegacyChildren(cmd *captain.Command, globals *globalOptions) {
	logging.Debug("register")

	secretsapi.InitializeClient()

	setLegacyOutput(globals)

	cmd.AddLegacyChildren(
		events.Command,
		update.Command,
		projects.Command,
		show.Command,
		run.Command,
		scripts.Command,
		pull.Command,
		export.Command,
		invite.Command,
		pkg.Command,
		secrets.NewCommand(secretsapi.Get(), &globals.Output).Config(),
		fork.Command,
	)
}
