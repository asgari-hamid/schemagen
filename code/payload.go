package code

type SchemaType int

const (
	SchemaTypeObject SchemaType = iota
	SchemaTypeArray
	SchemaTypeString
	SchemaTypeNumber
	SchemaTypeInteger
	SchemaTypeBoolean
)

type PayloadField struct {
	StructName string
	JsonName   string

	Type     SchemaType
	Nullable bool

	Tags map[string]string
}

type Payload struct {
	Name   string
	Fields []*PayloadField
}
