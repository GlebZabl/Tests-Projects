package requests

import "regexp"

type Request interface {
	Validate() bool
}

func checkRegExp(source string, exp string) bool {
	matched, err := regexp.MatchString(exp, source)
	if err != nil {
		return false
	}
	return matched
}
