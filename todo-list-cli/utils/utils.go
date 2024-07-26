package utils

import "strings"

func FormatDate(dateStringToFormat string) string {
	if idx := strings.LastIndex(dateStringToFormat, "."); idx != -1 {
		dateStringToFormat = dateStringToFormat[:idx] + "s"
	}
	return dateStringToFormat
}
