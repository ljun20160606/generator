package generator

import (
	"strings"
)

var dialect = []string{"ddl", "dml"}

func snakeToCamel(name string) string {
	//name = strings.ToLower(name)
	split := strings.Split(name, "_")
	builder := strings.Builder{}
SPLIT:
	for i := range split {
		unit := split[i]
		if len(unit) <= 0 {
			continue
		}
		if unit[0] >= 'A' && unit[0] <= 'Z' {
			builder.WriteString(unit)
			continue
		}
		for j := range dialect {
			v := dialect[j]
			lowerUnit := strings.ToLower(unit)
			if lowerUnit == v {
				builder.WriteString(strings.ToUpper(lowerUnit))
				builder.WriteString(unit[len(lowerUnit):])
				continue SPLIT
			}
		}
		builder.WriteString(strings.ToUpper(unit[:1]) + unit[1:])
	}
	return builder.String()
}

func snakeToCamelStartLower(name string) string {
	camel := snakeToCamel(name)
	for i := range dialect {
		s := dialect[i]
		if strings.HasPrefix(camel, strings.ToUpper(s)) {
			return strings.ToLower(s) + camel[len(s):]
		}
	}
	return strings.ToLower(camel[:1]) + camel[1:]
}
