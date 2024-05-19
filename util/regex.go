package util

import (
	"fmt"
	"regexp"
)

// RegexKeyValue is regex to parse key value error
// will return string which is contain about column and value
func RegexKeyValue(src, pattern string) (msg string, matches []string) {
	re := regexp.MustCompile(pattern)
	matches = re.FindStringSubmatch(src)

	if len(matches) > 2 {
		field := matches[1]
		value := matches[2]
		detailedMessage := fmt.Sprintf("The %s '%s'", field, value)
		return detailedMessage, matches
	}

	return "No details.", nil
}

func RegexKeyValueExist(src, pattern string, exist bool) (msg string, matches []string) {
	add := "already exists."
	if !exist {
		add = "is not exists."
	}

	msg, matches = RegexKeyValue(src, pattern)
	msg = msg + " " + add

	return
}
