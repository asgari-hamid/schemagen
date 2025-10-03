package gen

import (
	"github.com/asgari-hamid/schemagen/code"

	"github.com/dave/jennifer/jen"
)

func GeneratePayloadStruct(f *jen.File, p *code.Payload) {
	f.Type().Id(p.Name).StructFunc(func(g *jen.Group) {
		g.Id("Mask").Index().Bool()
		for _, field := range p.Fields {
			s := g.Id(field.StructName)
			switch field.Type {
			case code.SchemaTypeObject:
				panic("Not yet implemented")
			case code.SchemaTypeArray:
				panic("Not yet implemented")
			case code.SchemaTypeString:
				s = s.String()
			case code.SchemaTypeNumber:
				s = s.Float64()
			case code.SchemaTypeInteger:
				s = s.Int64()
			case code.SchemaTypeBoolean:
				s = s.Bool()
			default:
				panic("Unknown type")
			}
		}
	})
}

func GeneratePayloadFieldIndices(f *jen.File, p *code.Payload) {
	f.Const().DefsFunc(func(g *jen.Group) {
		for i, field := range p.Fields {
			if i == 0 {
				g.Id(p.Name + field.StructName + "Ref").Int().Op("=").Iota()
			} else {
				g.Id(p.Name + field.StructName + "Ref")
			}
			if i == len(p.Fields)-1 {
				g.Id(p.Name + "FieldCount")
			}
		}
	})
}

func GeneratePayloadFieldMask(f *jen.File, p *code.Payload) {
	f.Func().
		Id("Build" + p.Name + "FieldMask").
		Params(jen.Id("fields").Index().String()).
		Index().Bool().
		BlockFunc(func(g *jen.Group) {
			g.Id("mask").Op(":=").Make(jen.Index().Bool(), jen.Id(p.Name+"FieldCount"))

			g.For(
				jen.List(jen.Id("_"), jen.Id("field")).Op(":=").Range().Id("fields"),
			).BlockFunc(func(g *jen.Group) {
				g.Switch(jen.Id("field")).BlockFunc(func(g *jen.Group) {
					for _, field := range p.Fields {
						g.Case(jen.Lit(field.JsonName)).
							Id("mask").Index(jen.Id(p.Name + field.StructName + "Ref")).Op("=").True()
					}
				})
			})

			g.Return(jen.Id("mask"))
		})
}

func GeneratePayloadJsonWriter(f *jen.File, p *code.Payload) {
	f.Func().
		Params(jen.Id("x").Op("*").Id(p.Name)).
		Id("writeJson").
		Params(
			jen.Id("writer").Op("*").Qual("github.com/asgari-hamid/jsonw", "ObjectWriter"),
			jen.Id("mask").Index().Bool(),
		).
		BlockFunc(func(g *jen.Group) {
			g.Id("writer").Dot("Open").Call()
			for _, field := range p.Fields {
				g.If(
					jen.Id("mask").Index(jen.Id(p.Name + field.StructName + "Ref")),
				).BlockFunc(func(g *jen.Group) {
					method := ""
					switch field.Type {
					case code.SchemaTypeObject:
						panic("Not yet implemented")
					case code.SchemaTypeArray:
						panic("Not yet implemented")
					case code.SchemaTypeString:
						method = "StringField"
					case code.SchemaTypeNumber:
						method = "FloatField"
					case code.SchemaTypeInteger:
						method = "IntegerField"
					case code.SchemaTypeBoolean:
						method = "BooleanField"
					default:
						panic("Unknown type")
					}
					g.Id("writer").Dot(method).Call(jen.Lit(field.JsonName), jen.Id("x").Dot(field.StructName))
				})
			}
			g.Id("writer").Dot("Close").Call()
		})
}

func GeneratePayloadMarshaler(f *jen.File, p *code.Payload) {
	f.Func().
		Params(jen.Id("x").Op("*").Id(p.Name)).
		Id("MarshalJSON").
		Params().
		Params(jen.Index().Byte(), jen.Error()).
		BlockFunc(func(g *jen.Group) {
			g.Id("writer").Op(":=").Qual("github.com/asgari-hamid/jsonw", "NewObjectWriter").Call(jen.Nil())
			g.Id("x").Dot("writeJson").Call(jen.Id("writer"), jen.Id("x").Dot("Mask"))
			g.Return(jen.Id("writer").Dot("BuildBytes").Call())
		})
}
