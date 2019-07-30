package generator

import (
	"strconv"
	"strings"
)

type Config struct {
	Dest          string        `json:"dest"`
	Pkg           string        `json:"pkg"`
	ConnectConfig ConnectConfig `json:"connectConfig"`
	ModelConfig   ModelConfig   `json:"modelConfig"`
}

type ConnectConfig struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Scheme   string `json:"scheme"`
}

// username:password@tcp(mysql.com:3317)/test
func (s ConnectConfig) String() string {
	builder := strings.Builder{}
	if s.Username != "" && s.Password != "" {
		builder.WriteString(s.Username)
		builder.WriteString(":")
		builder.WriteString(s.Password)
		builder.WriteString("@")
	}
	builder.WriteString("tcp(")
	builder.WriteString(s.Host)
	builder.WriteString(":")
	builder.WriteString(strconv.Itoa(s.Port))
	builder.WriteString(")/")
	builder.WriteString(s.Scheme)
	return builder.String()
}

type TableConfig struct {
	Includes []string `json:"includes"`
	Excludes []string `json:"excludes"`
}

type ModelConfig struct {
	Filename      string      `json:"filename"`
	FileExtension string      `json:"fileExtension"`
	TableConfig   TableConfig `json:"tableConfig"`
}

// return "a" if filename is "./a"
func (m *ModelConfig) name() string {
	split := strings.Split(m.Filename, "/")
	return split[len(split)-1]
}
