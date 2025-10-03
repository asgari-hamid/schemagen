package gen

import (
	"schemagen/code"

	"github.com/dave/jennifer/jen"
)

func GeneratePayloadStruct(j *jen.File, p *code.Payload) {
	j.Type().Id(p.Name).StructFunc(func(g *jen.Group) {
		for _, field := range p.Fields {
			g.Id(field.Name).Id(field.Type).Tag(field.Tags)
		}
	})
}
