package validators

import (
	"regexp"
	"strings"
)

var emailRegex = regexp.MustCompile(
	`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`,
)

func IsEmailValid(email string) bool {
	email = strings.TrimSpace(email)

	if len(email) < 5 || len(email) > 254 {
		return false
	}

	return emailRegex.MatchString(email)
}
