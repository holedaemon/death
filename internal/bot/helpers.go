package bot

import (
	"strings"

	"github.com/diamondburned/arikawa/v3/discord"
)

func pluralize(s string) string {
	if strings.HasSuffix(s, "s") {
		return s + "'"
	} else {
		return s + "'s"
	}
}

func roleInSlice(role discord.RoleID, roles []discord.RoleID) bool {
	for _, r := range roles {
		if role == r {
			return true
		}
	}

	return false
}
