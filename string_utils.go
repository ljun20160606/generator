package generator

import (
	"regexp"
	"strings"
)

func snakeToCamel(name string) string {
	name = strings.ToLower(name)
	name = regexp.MustCompile(`_[a-zA-Z]`).ReplaceAllStringFunc(name, func(from string) string {
		return strings.ToUpper(from[1:])
	})
	return name
}

func snakeToCamelFirstUpper(name string) string {
	camel := snakeToCamel(name)
	return strings.ToUpper(camel[:1]) + camel[1:]
}
