package auths

import "strings"

func splitToken(token string) []string {
	return strings.Split(token, ".")
}
