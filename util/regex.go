package util

import (
	"fmt"
	"regexp"

	"github.com/umardev500/banksampah/constant"
)

// RegexDuplicate is regex to parse duplicate entry
// will return string which is contain about column and value
func RegexDuplicate(src string) (msg string, matches []string) {
	re := regexp.MustCompile(string(constant.SqlErrPatternDuplicate))
	matches = re.FindStringSubmatch(src)

	if len(matches) > 2 {
		field := matches[1]
		value := matches[2]
		detailedMessage := fmt.Sprintf("%s %s already exists.", field, value)
		return detailedMessage, matches
	}

	return "No details.", nil
}
