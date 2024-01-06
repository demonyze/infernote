package utils

import "os"

func GetArg(index int, fallback string) string {
	args := os.Args[1:]

	if index >= len(args) {
		return fallback
	}

	var value string = args[index]

	if args[index] == "" {
		return fallback
	}

	return value
}
