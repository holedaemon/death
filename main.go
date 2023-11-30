package main

import (
	"context"
	"fmt"
	"os"

	"github.com/caarlos0/env/v7"
	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/holedaemon/death/internal/bot"
	"github.com/zikaeroh/ctxlog"
	"go.uber.org/zap"
)

type options struct {
	Token     string `env:"DEATH_TOKEN"`
	GuildID   string `env:"DEATH_GUILD_ID"`
	ChannelID string `env:"DEATH_CHANNEL_ID"`
	UserID    string `env:"DEATH_USER_ID"`
	RoleID    string `env:"DEATH_ROLE_ID"`
	RoleColor int32  `env:"DEATH_ROLE_COLOR"`
}

func main() {
	opts := &options{}
	eo := env.Options{
		RequiredIfNoDef: true,
	}

	if err := env.Parse(opts, eo); err != nil {
		fmt.Fprintf(os.Stderr, "error parsing environment variables into struct: %s\n", err.Error())
		return
	}

	logger := ctxlog.New(false)

	guildID, err := discord.ParseSnowflake(opts.GuildID)
	if err != nil {
		logger.Fatal("error parsing guild ID to snowflake", zap.Error(err))
	}

	channelID, err := discord.ParseSnowflake(opts.ChannelID)
	if err != nil {
		logger.Fatal("error parsing channel ID to snowflake", zap.Error(err))
	}

	userID, err := discord.ParseSnowflake(opts.UserID)
	if err != nil {
		logger.Fatal("error parsing user ID to snowflake", zap.Error(err))
	}

	roleID, err := discord.ParseSnowflake(opts.RoleID)
	if err != nil {
		logger.Fatal("error parsing role ID to snowflake", zap.Error(err))
	}

	b, err := bot.New(
		opts.Token,
		bot.WithLogger(logger),
		bot.WithGuildID(discord.GuildID(guildID)),
		bot.WithChannelID(discord.ChannelID(channelID)),
		bot.WithUserID(discord.UserID(userID)),
		bot.WithRoleID(discord.RoleID(roleID)),
		bot.WithRoleColor(discord.Color(opts.RoleColor)),
	)

	if err != nil {
		logger.Fatal("error creating bot", zap.Error(err))
	}

	ctx := context.Background()
	if err := b.Start(ctx); err != nil {
		logger.Fatal("error starting bot", zap.Error(err))
	}
}
