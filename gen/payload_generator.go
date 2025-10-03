package gen

import (
	"github.com/asgari-hamid/schemagen/code"

	"github.com/dave/jennifer/jen"
)

func GenerateImports(f *jen.File) {
	f.ImportNames(map[string]string{
		moPath:    "mo",
		jsonwPath: "jsonw",
	})
}

func GeneratePayloadStruct(f *jen.File, p *code.Payload) {
	f.Type().Id(p.Name).StructFunc(func(g *jen.Group) {
		g.Id("Mask").Index().Bool()

		for _, field := range p.Fields {
			addTypeDefinition(g.Id(field.StructName), field.Type, field.Nullable)
		}
	})
}

func GeneratePayloadFieldIndices(f *jen.File, p *code.Payload) {
	f.Const().DefsFunc(func(g *jen.Group) {
		for i, field := range p.Fields {
			if i == 0 {
				g.Id(getFieldRefName(p, field)).Int().Op("=").Iota()
			} else {
				g.Id(getFieldRefName(p, field))
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
							Id("mask").Index(jen.Id(getFieldRefName(p, field))).Op("=").True()
					}
				})
			})

			g.Return(jen.Id("mask"))
		})
}

func GeneratePayloadJsonWriter(f *jen.File, p *code.Payload) {
	f.Func().
		Params(jen.Id("x").Op("*").Id(p.Name)).
		Id("WriteJson").
		Params(jen.Id("writer").Op("*").Qual(jsonwPath, "ObjectWriter")).
		BlockFunc(func(g *jen.Group) {
			g.Id("mask").Op(":=").Id("x").Dot("Mask")
			g.Id("noMask").Op(":=").Len(jen.Id("mask")).Op("!=").Id(p.Name + "FieldCount")
			g.Id("writer").Dot("Open").Call()
			for _, field := range p.Fields {
				g.If(
					jen.Id("noMask").Op("||").Id("mask").Index(jen.Id(getFieldRefName(p, field))),
				).BlockFunc(func(g *jen.Group) {
					if field.Nullable {
						g.If(
							jen.List(jen.Id("value"), jen.Id("exists")).Op(":=").Id("x").Dot(field.StructName).Dot("Get").Call().Op(";").Id("exists"),
						).BlockFunc(func(g *jen.Group) {
							method := mapSchemaType(field.Type)
							g.Id("writer").Dot(method).Call(jen.Lit(field.JsonName), jen.Id("value"))
						}).Else().BlockFunc(func(g *jen.Group) {
							g.Id("writer").Dot("NullField").Call(jen.Lit(field.JsonName))
						})
					} else {
						method := mapSchemaType(field.Type)
						g.Id("writer").Dot(method).Call(jen.Lit(field.JsonName), jen.Id("x").Dot(field.StructName))
					}
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
			g.Id("writer").Op(":=").Qual(jsonwPath, "NewObjectWriter").Call(jen.Nil())
			g.Id("x").Dot("WriteJson").Call(jen.Id("writer"))
			g.Return(jen.Id("writer").Dot("BuildBytes").Call())
		})
}
