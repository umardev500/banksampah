package util

import (
	"regexp"
	"strconv"
	"strings"
)

func StrToInt(str string, defaultValue int) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		return defaultValue
	}

	return i
}

func RemoveExtraSpace(str string) string {
	re := regexp.MustCompile(`\s+`)

	result := re.ReplaceAllString(str, " ")
	return result
}

func StringTrimAnNoExtraSpace(str string) string {
	return strings.TrimSpace(RemoveExtraSpace(str))
}

func RemoveSqlComment(str string) string {
	re := regexp.MustCompile(`--.*`)
	return re.ReplaceAllString(str, "")
}
