package utils

func AlternateSuffix(suffix string) string {
	switch suffix {
	case "major":
		return ""

	case "minor":
		return "m"

	default:
		return suffix
	}
}
