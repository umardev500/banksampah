package util

import (
	"regexp"
	"strings"
)

// RegexKeyValue is regex to parse key value error
// will return string which is contain about column and value
func RegexKeyValue(src, pattern string) (matches []string) {
	re := regexp.MustCompile(pattern)
	matches = re.FindStringSubmatch(src)

	if len(matches) > 2 {
		return matches
	}

	return nil
}

func removeExtraMessage(src, prefix string) (result string) {
	ptrn := `from table ".*?".`
	re := regexp.MustCompile(ptrn)
	result = re.ReplaceAllString(src, "")
	result = strings.TrimSpace(strings.TrimPrefix(result, prefix))
	matched := re.MatchString(src)
	if matched {
		result += "."
	}
	return
}

func RegexKeyValueExist(src, pattern string, exist bool) (msg string, matches []string) {
	matches = RegexKeyValue(src, pattern)
	msg = removeExtraMessage(src, matches[0])

	return
}
