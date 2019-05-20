package generator

import "database/sql"

type Table struct {
	Name    string `db:"name"`
	Comment string `db:"table_comment"`
}

type Column struct {
	Field      string         `db:"Field"`
	Type       string         `db:"Type"`
	Null       sql.NullString `db:"Null"`
	Key        []byte         `db:"Key"`
	Default    []byte         `db:"Default"`
	Comment    string         `db:"Comment"`
	Extra      []byte         `db:"Extra"`
	Privileges []byte         `db:"Privileges"`
	Collation  []byte         `db:"Collation"`
}
