package bot

import (
	"time"

	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/diamondburned/arikawa/v3/gateway"
	"go.uber.org/zap"
)

func (b *Bot) onReady(r *gateway.ReadyEvent) {
	b.logger.Info("connected to discord", zap.String("user_id", r.User.ID.String()))
}

func (b *Bot) onGuildMemberAdd(g *gateway.GuildMemberAddEvent) {
	if g.GuildID != b.guildID {
		return
	}

	if g.User.ID != b.userID {
		return
	}

	b.logger.Info("user has joined the server")

	guild, err := b.state.Guild(b.guildID)
	if err != nil {
		b.logger.Error("error fetching guild, can't apply roles", zap.Error(err))
		return
	}

	ignore := make([]discord.RoleID, 0)
	for _, r := range guild.Roles {
		if r.Color.String() == b.roleColor.String() {
			if err := b.state.AddRole(b.guildID, b.userID, r.ID, api.AddRoleData{
				AuditLogReason: "Re-adding role after ban.",
			}); err != nil {
				b.logger.Error("error adding role to user", zap.Error(err), zap.String("role_id", r.ID.String()))
			}

			ignore = append(ignore, r.ID)
			time.Sleep(time.Second * 2)
		}
	}

	member, err := b.state.Member(g.GuildID, g.User.ID)
	if err != nil {
		b.logger.Error("error fetching member", zap.Error(err))
		return
	}

	for _, r := range member.RoleIDs {
		if roleInSlice(r, ignore) {
			continue
		}

		if err := b.state.RemoveRole(g.GuildID, g.User.ID, r, "It's not the right color!!!"); err != nil {
			b.logger.Error("error removing role from user", zap.Error(err), zap.String("role_id", r.String()))
		}
	}
}
