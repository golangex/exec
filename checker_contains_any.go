package gex

import (
	`strings`
)

type containsAnyChecker struct {
	items []string

	cache bool
	all   strings.Builder
}

func (cac *containsAnyChecker) check(line string) (checked bool, err error) {
	if cac.cache {
		cac.all.WriteString(line)
	}

	checked = cac.containsAny(line, cac.items...)
	if !checked && cac.cache {
		checked = cac.containsAny(cac.all.String(), cac.items...)
	}

	return
}

func (cac *containsAnyChecker) containsAny(content string, items ...string) (contains bool) {
	for _, item := range items {
		contains = strings.Contains(content, item)
		if contains {
			break
		}
	}

	return
}
