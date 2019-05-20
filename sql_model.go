package generator

import "strings"

type SqlModel struct {
	Package   string
	Comment   string
	Imports   map[interface{}]string
	ClassName string
	Cols      []RealColumn
}

func (j *SqlModel) parseMysqlColumn(column Column) {
	col := RealColumn{}
	col.Column = column
	col.CanNull = column.Null.String != "NO"
	col.DBName = column.Field
	col.Name = snakeToCamel(column.Field)
	col.Comment = strings.Replace(column.Comment, "\n", `\n *`, -1)
	colType := column.Type

	switch {
	case strings.Contains(colType, "char") || strings.Contains(colType, "text"):
		setJavaType(&col, j, JavaString)
	case strings.Contains(colType, "bigint"):
		setJavaType(&col, j, JavaLong)
	case strings.Contains(colType, "tinyint") ||
		strings.Contains(colType, "bool") ||
		strings.Contains(colType, "int"):
		setJavaType(&col, j, JavaInteger)
	case strings.Contains(colType, "float"):
		setJavaType(&col, j, JavaFloat)
	case strings.Contains(colType, "double"):
		setJavaType(&col, j, JavaDouble)
	case strings.Contains(colType, "time") ||
		strings.Contains(colType, "date") ||
		strings.Contains(colType, "datetime") ||
		strings.Contains(colType, "timestamp"):
		setJavaType(&col, j, JavaDate)
	case strings.Contains(colType, "blob"):
		setJavaType(&col, j, JavaBytes)
	}
	j.Cols = append(j.Cols, col)
}

func setJavaType(col *RealColumn, model *SqlModel, jt javaType) {
	col.Type = jt.String()
	im, has := javaImports[jt]
	if !has {
		return
	}
	model.Imports[jt] = im
}

type RealColumn struct {
	Name    string
	Type    string
	DBName  string
	CanNull bool
	Comment string
	Column  Column
}
