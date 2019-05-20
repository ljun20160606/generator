package generator

type javaType int

const (
	JavaString javaType = iota
	JavaLong
	JavaInteger
	JavaFloat
	JavaDouble
	JavaDate
	JavaBytes
	JavaByte
)

var javaImports = map[javaType]string{
	JavaDate: "java.util.Date",
}

var javaTypeStrings = [...]string{
	JavaString:  "String",
	JavaLong:    "Long",
	JavaInteger: "Integer",
	JavaFloat:   "Float",
	JavaDouble:  "Double",
	JavaDate:    "Date",
	JavaBytes:   "Byte[]",
	JavaByte:    "Byte",
}

func (j javaType) String() string {
	return javaTypeStrings[j]
}
