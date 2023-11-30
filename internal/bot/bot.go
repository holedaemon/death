package bot

import (
	"context"
	"fmt"

	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/diamondburned/arikawa/v3/gateway"
	"github.com/diamondburned/arikawa/v3/state"
	"github.com/zikaeroh/ctxlog"
	"go.uber.org/zap"
)

// Bot is a Discord client responsible for updating Max's role and occasionally
// banning her.
type Bot struct {
	state  *state.State
	logger *zap.Logger

	userID    discord.UserID
	channelID discord.ChannelID
	guildID   discord.GuildID
	roleColor discord.Color
	roleID    discord.RoleID

	firstCycle bool
}

// New creates a new Bot.
func New(token string, opts ...Option) (*Bot, error) {
	if token == "" {
		return nil, fmt.Errorf("bot: token cannot be blank")
	}

	b := &Bot{
		firstCycle: true,
	}

	for _, o := range opts {
		o(b)
	}

	if b.logger == nil {
		b.logger = ctxlog.New(false)
	}

	if b.channelID == 0 {
		return nil, fmt.Errorf("bot: missing channel ID")
	}

	if b.guildID == 0 {
		return nil, fmt.Errorf("bot: missing guild ID")
	}

	if b.roleID == 0 {
		return nil, fmt.Errorf("bot: missing role ID")
	}

	if b.userID == 0 {
		return nil, fmt.Errorf("bot: missing user ID")
	}

	if b.roleColor == 0 {
		return nil, fmt.Errorf("bot: missing role color")
	}

	b.state = state.New("Bot " + token)
	b.state.AddIntents(gateway.IntentGuilds | gateway.IntentGuildMembers)
	b.state.AddHandler(b.onReady)
	b.state.AddHandler(b.onGuildMemberAdd)

	return b, nil
}

// Start begins necessary goroutines and connects to Discord.
func (b *Bot) Start(ctx context.Context) error {
	ctx = ctxlog.WithLogger(ctx, b.logger)
	go b.updater(ctx)

	return b.state.Connect(ctx)
}
