package main

import (
	"github.com/asgari-hamid/schemagen/code"
	"github.com/asgari-hamid/schemagen/gen"

	"github.com/dave/jennifer/jen"
)

func main() {
	j := jen.NewFile("payloads")

	p := &code.Payload{
		Name: "Product",
		Fields: []*code.PayloadField{
			{StructName: "Id", JsonName: "id", Type: code.SchemaTypeString, Tags: nil},
			{StructName: "Title", JsonName: "title", Type: code.SchemaTypeString, Tags: nil},
			{StructName: "Description", JsonName: "description", Type: code.SchemaTypeString, Tags: nil},
			{StructName: "Price", JsonName: "price", Type: code.SchemaTypeNumber, Tags: nil},
			{StructName: "Weight", JsonName: "weight", Type: code.SchemaTypeInteger, Tags: nil},
			{StructName: "Published", JsonName: "published", Type: code.SchemaTypeBoolean, Tags: nil},
		},
	}

	gen.GeneratePayloadFieldIndices(j, p)
	gen.GeneratePayloadFieldMask(j, p)
	gen.GeneratePayloadStruct(j, p)
	gen.GeneratePayloadJsonWriter(j, p)

	if err := j.Save("./payloads/product_gen.go"); err != nil {
		panic(err)
	}
}
