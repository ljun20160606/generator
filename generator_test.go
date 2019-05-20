package generator

import (
	"testing"
)

func TestName(t *testing.T) {
	gen := new(Generator)
	gen.rawGen(&GeneratorConfig{
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
