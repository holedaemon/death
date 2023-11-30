package bot

import "strings"

func pluralize(s string) string {
	if strings.HasSuffix(s, "s") {
		return s + "'"
	} else {
		return s + "'s"
	}
}
