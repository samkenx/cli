package cmdtree

import (
	"github.com/ActiveState/cli/internal/prompt"
	"github.com/ActiveState/cli/internal/runners/invite"
	"github.com/ActiveState/cli/pkg/project"

	"github.com/ActiveState/cli/internal/captain"
	"github.com/ActiveState/cli/internal/locale"
	"github.com/ActiveState/cli/internal/output"
)

func newInviteCommand(pj *project.Project, out output.Outputer, prompt prompt.Prompter) *captain.Command {
	inviteRunner := invite.New(pj, out, prompt)

	params := invite.Params{}

	return captain.NewCommand(
		"invite",
		locale.Tl("invite_description", "Invite new members to an organization"),
		[]*captain.Flag{
			{
				Name:        "organization",
				Description: locale.Tl("invite_flag_organization_description", "Organization to invite to. If not set, invite to current project's organization"),
				Value:       &params.Org,
			},
			{
				Name:        "role",
				Description: locale.Tl("invite_flag_role_description", "Set user role to 'member' or 'owner'. If not set, prompt for the role"),
				Value:       &params.Role,
			},
		},
		[]*captain.Argument{
			&captain.Argument{
				Name:        "email1,[email2,..]",
				Description: locale.Tl("invite_arg_emails", "Email addresses to send the invitations to"),
				Required:    true,
				Value:       &params.EmailList,
			},
		},
		func(ccmd *captain.Command, _ []string) error {
			return inviteRunner.Run(&params)
		},
	)
}
