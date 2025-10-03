package gen

import (
	"github.com/asgari-hamid/schemagen/code"
	"github.com/dave/jennifer/jen"
)

const (
	moPath    = "github.com/samber/mo"
	jsonwPath = "github.com/asgari-hamid/jsonw"
)

func getFieldRefName(p *code.Payload, f *code.PayloadField) string {
	return p.Name + f.StructName + "Ref"
}

func addTypeDefinition(s *jen.Statement, t code.SchemaType, nullable bool) {
	switch t {
	case code.SchemaTypeObject:
		panic("Not yet implemented")
	case code.SchemaTypeArray:
		panic("Not yet implemented")
	case code.SchemaTypeString:
		if nullable {
			s.Qual(moPath, "Option").Types(jen.String())
		} else {
			s.String()
		}
	case code.SchemaTypeNumber:
		if nullable {
			s.Qual(moPath, "Option").Types(jen.Float64())
		} else {
			s = s.Float64()
		}
	case code.SchemaTypeInteger:
		if nullable {
			s.Qual(moPath, "Option").Types(jen.Int64())
		} else {
			s = s.Int64()
		}
	case code.SchemaTypeBoolean:
		if nullable {
			s.Qual(moPath, "Option").Types(jen.Bool())
		} else {
			s = s.Bool()
		}
	default:
		panic("Unknown schema type")
	}
}

func mapSchemaType(t code.SchemaType) string {
	method := ""
	switch t {
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
		panic("Unknown schema type")
	}
	return method
}
