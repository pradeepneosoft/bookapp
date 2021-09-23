package helper

import "strings"

func GetTokenFromHeader(token string) string {
	return strings.Split(token, " ")[1]
}
