package generator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestName(t *testing.T) {
	gen := new(Generator)
	gen.rawGen(&Config{
		Dest: "./model",
		Pkg:  "com.ljun20160606.mysql",
		ConnectConfig: ConnectConfig{
			Username: "test",
			Password: "1234566",
			Host:     "mysql.com",
			Port:     3317,
			Scheme:   "test",
		},
	})
}

func TestFilterTables(t *testing.T) {
	ast := assert.New(t)

	tables := []Table{
		{"Foo", ""},
		{"Bar", ""},
		{"Fooz", ""},
		{"Barz", ""},
	}

	// no rule
	ast.Equal(tables, filterTables(tables, []string{}, []string{}))

	// include Foo
	ast.Equal([]Table{{"Foo", ""}}, filterTables(tables, []string{"Foo"}, []string{}))

	// include Foo, exclude Foo
	ast.Equal(len([]Table{}), len(filterTables(tables, []string{"Foo"}, []string{"Foo", "Bar"})))

	// include Foo Bar, exclude Foo
	ast.Equal([]Table{{"Bar", ""}}, filterTables(tables, []string{"Foo", "Bar"}, []string{"Foo"}))
}
