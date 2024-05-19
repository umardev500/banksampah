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
		value := matches[2]
		detailedMessage := fmt.Sprintf("%s already exists.", value)
		return detailedMessage, matches
	}

	return "No details.", nil
}
