package bot

import (
	"math/rand"
	"strings"
)

var messages = []string{
	"%USER% was found dead in a ditch",
	"%USER% has passed on",
	"%USER% is on their way to see King Kai",
	"%USER% got acute radiation poisoning",
	"%USER% fell from a high place",
	"%USER% has met a terrible fate",
	"%USER% was devoured by Turfy's cat",
}

func (b *Bot) sendMessage() error {
	user, err := b.state.User(b.userID)
	if err != nil {
		return err
	}

	msg := messages[rand.Intn(len(messages))]
	msg = strings.Replace(msg, "%USER%", user.DisplayOrUsername(), -1)

	_, err = b.state.SendMessage(b.channelID, msg)
	return err
}
