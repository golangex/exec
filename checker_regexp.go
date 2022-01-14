package gex

import (
	`regexp`
	`strings`
)

type regexpChecker struct {
	regexp string

	cache bool
	all   strings.Builder
}

func (rc *regexpChecker) check(line string) (checked bool, err error) {
	if rc.cache {
		rc.all.WriteString(line)
	}

	if checked, err = regexp.MatchString(rc.regexp, line); nil != err {
		return
	}
	if !checked && rc.cache {
		checked, err = regexp.MatchString(rc.regexp, rc.all.String())
	}

	return
}
