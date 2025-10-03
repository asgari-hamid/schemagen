package code

type PayloadField struct {
	Name string
	Type string
	Tags map[string]string
}

type Payload struct {
	Name   string
	Fields []*PayloadField
}
