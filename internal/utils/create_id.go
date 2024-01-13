package utils

import (
	"strings"
)

func CreateId(key string, suffix string) string {
	formattedKey := strings.ReplaceAll(key, " ", "")
	formattedSuffix := strings.ReplaceAll(suffix, " ", "")

	return strings.ToLower(formattedKey + formattedSuffix)
}
