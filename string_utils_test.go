package generator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSnakeToCamel(t *testing.T) {
	ast := assert.New(t)
	ast.Equal(snakeToCamel("ddl"), "DDL")
	ast.Equal(snakeToCamel("ddl_"), "DDL")
	ast.Equal(snakeToCamel("_ddl"), "DDL")
	ast.Equal(snakeToCamel("dml"), "DML")
	ast.Equal(snakeToCamel("dml_"), "DML")
	ast.Equal(snakeToCamel("_dml"), "DML")
	ast.Equal(snakeToCamel("ddl_operate"), "DDLOperate")
	ast.Equal(snakeToCamel("foo"), "Foo")
	ast.Equal(snakeToCamel("fFo"), "FFo")
	ast.Equal(snakeToCamel("foo_bar"), "FooBar")
	ast.Equal(snakeToCamel("foo_bar_zoo"), "FooBarZoo")
}
