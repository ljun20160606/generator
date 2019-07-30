package main

import (
	"github.com/ljun20160606/generator"
	"github.com/ljun20160606/gox/fs"
)

func main() {
	config := new(generator.Config)
	err := fs.ReadJSON("./generator.json", config)
	if err != nil {
		panic(err)
	}
	generator.Gen(config)
}
