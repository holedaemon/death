package bot

import (
	"github.com/diamondburned/arikawa/v3/discord"
	"go.uber.org/zap"
)

// Option configures a bot.
type Option func(*Bot)

// WithLogger sets a bot's logger.
func WithLogger(l *zap.Logger) Option {
	return func(b *Bot) {
		b.logger = l
	}
}

// WithUserID sets a bot's user ID.
func WithUserID(uid discord.UserID) Option {
	return func(b *Bot) {
		b.userID = uid
	}
}

// WithRoleID sets a bot's role ID.
func WithRoleID(rid discord.RoleID) Option {
	return func(b *Bot) {
		b.roleID = rid
	}
}

// WithGuildID sets a bot's guild ID.
func WithGuildID(gid discord.GuildID) Option {
	return func(b *Bot) {
		b.guildID = gid
	}
}

// WithChannelID sets a bot's channel ID.
func WithChannelID(cid discord.ChannelID) Option {
	return func(b *Bot) {
		b.channelID = cid
	}
}

// WithRoleColor sets a bot's role color.
func WithRoleColor(c discord.Color) Option {
	return func(b *Bot) {
		b.roleColor = c
	}
}
