package bot

import (
	"context"
	"math/rand"
	"strings"
	"time"

	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/utils/json/option"
	"github.com/zikaeroh/ctxlog"
	"go.uber.org/zap"
)

func (b *Bot) updater(ctx context.Context) {
	// t := time.NewTicker(time.Hour * (24 * 7))
	t := time.NewTicker(time.Second * 10)

	for {
		select {
		case <-ctx.Done():
			t.Stop()
			return
		case <-t.C:
			if rand.Intn(100) == 1 {
				if b.firstCycle {
					b.firstCycle = false
					continue
				}

				var reason string
				user, err := b.state.User(b.userID)
				if err != nil {
					ctxlog.Error(ctx, "error fetching user", zap.Error(err))
					reason = "Death strikes quick"
				} else {
					reason = "Death strikes quick; " + pluralize(user.DisplayOrUsername()) + " time has come"
				}

				if err := b.state.Ban(b.guildID, b.userID, api.BanData{
					DeleteDays:     option.NewUint(0),
					AuditLogReason: api.AuditLogReason(reason),
				}); err != nil {
					ctxlog.Error(ctx, "error banning member from guild", zap.Error(err))
					continue
				}

				err = b.sendMessage()
				if err != nil {
					ctxlog.Error(ctx, "error sending message", zap.Error(err))
					continue
				}

				continue
			}

			if rand.Intn(2) == 1 {
				role, err := b.state.Role(b.guildID, b.roleID)
				if err != nil {
					ctxlog.Error(ctx, "error fetching role", zap.Error(err))
					continue
				}

				newName := role.Name
				name := strings.ToLower(role.Name)
				if name == "alive" {
					newName = "dead"
				} else if name == "dead" {
					newName = "alive"
				}

				if newName == role.Name {
					ctxlog.Info(ctx, "unexpected role name, skipping...")
					continue
				}

				if _, err := b.state.ModifyRole(b.guildID, b.roleID, api.ModifyRoleData{
					Name: option.NewNullableString(newName),
				}); err != nil {
					ctxlog.Error(ctx, "error modifying role name", zap.Error(err))
					continue
				}
				continue
			}

			ctxlog.Info(ctx, "neither conditions hit, see you next time")
		}
	}
}
