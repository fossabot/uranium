package validate

import (
	"regexp"

	"github.com/go-uranium/uranium/utils/clean"
)

var UsernameRegex = regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9_]{0,19}$`)
var UsernameMax = 20

func Username(u string) (bool, string) {
	if len(u) < 1 || len(u) > UsernameMax {
		return false, ""
	}
	u = clean.Username(u)
	if !UsernameRegex.MatchString(u) {
		return false, ""
	}
	return true, u
}
