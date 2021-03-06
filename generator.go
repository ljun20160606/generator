package generator

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/ljun20160606/gox/fs"
	"io"
	"log"
	"strings"
	"text/template"
)

func Gen(config *Config) {
	gen := new(Generator)
	gen.rawGen(config)
}

type Generator struct {
	db              *sqlx.DB
	specifiedTables []string
}

const selectTables string = `SELECT
 table_name AS name,table_comment AS table_comment
FROM
 information_schema.tables
WHERE
 table_schema = DATABASE()`

func (g *Generator) rawGen(config *Config) {
	t, err := template.New(config.ModelConfig.name()).ParseFiles(config.ModelConfig.Filename)
	if err != nil {
		log.Fatal(err)
	}

	g.db = sqlx.MustOpen("mysql", config.ConnectConfig.String())
	defer g.db.Close()

	var tables []Table
	err = g.db.Select(&tables, selectTables)
	if err != nil {
		log.Fatal(err)
	}

	tables = filterTables(tables, config.ModelConfig.TableConfig.Includes, config.ModelConfig.TableConfig.Excludes)

	for i := range tables {
		g.genTable(config.Dest, config.Pkg, config.ModelConfig.FileExtension, t, tables[i])
	}
}

func filterTables(tables []Table, includes, excludes []string) []Table {
	var includedTables []Table
	if len(includes) != 0 {
		contains := func(value string) int {
			for i := range tables {
				if tables[i].Name == value {
					return i
				}
			}
			return -1
		}
		for i := range includes {
			if index := contains(includes[i]); index != -1 {
				includedTables = append(includedTables, tables[index])
			}
		}
	} else {
		includedTables = tables
	}

	var excludedTables []Table
	if len(excludes) != 0 {
		contains := func(value string) int {
			for i := range excludes {
				if excludes[i] == value {
					return i
				}
			}
			return -1
		}
		for i := range includedTables {
			if index := contains(includedTables[i].Name); index == -1 {
				excludedTables = append(excludedTables, includedTables[i])
			}
		}
	} else {
		excludedTables = includedTables
	}
	return excludedTables
}

func (g *Generator) genTable(dest, pkg, fileExtension string, t *template.Template, table Table) {
	var cols []Column
	err := g.db.Select(&cols, "SHOW FULL COLUMNS FROM `"+table.Name+"`")
	if err != nil {
		panic(err)
	}
	model := SqlModel{
		Package:   pkg,
		Comment:   strings.Replace(table.Comment, "\n", "\n *", -1),
		Imports:   make(map[interface{}]string),
		ClassName: snakeToCamel(table.Name),
	}

	for i := range cols {
		model.parseMysqlColumn(cols[i])
	}
	_ = fs.MkdirP(dest)
	err = fs.WriteFile(dest+"/"+model.ClassName+"."+fileExtension, func(writer io.Writer) error {
		err = t.Execute(writer, model)
		if err != nil {
			log.Fatal(err)
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}
