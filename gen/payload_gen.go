package gen

import (
	"schemagen/code"

	"github.com/dave/jennifer/jen"
)

func GeneratePayloadStruct(f *jen.File, p *code.Payload) {
	f.Type().Id(p.Name).StructFunc(func(g *jen.Group) {
		for _, field := range p.Fields {
			g.Id(field.Name).Id(field.Type).Tag(field.Tags)
		}
	})
}

func GeneratePayloadFieldIndices(f *jen.File, p *code.Payload) {
	f.Const().DefsFunc(func(g *jen.Group) {
		for i, field := range p.Fields {
			if i == 0 {
				g.Id(p.Name + field.Name + "Ref").Int().Op("=").Iota()
			} else {
				g.Id(p.Name + field.Name + "Ref")
			}
			if i == len(p.Fields)-1 {
				g.Id(p.Name + "FieldCount")
			}
		}
	})
}
