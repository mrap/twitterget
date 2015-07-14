package user

import "strings"

func ToUsername(username string) string {
	if strings.HasPrefix(username, "@") {
		return username
	} else {
		return "@" + username
	}
}
