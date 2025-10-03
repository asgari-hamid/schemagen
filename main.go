package main

import (
	"encoding/json"
	"fmt"

	"github.com/asgari-hamid/schemagen/code"
	"github.com/asgari-hamid/schemagen/gen"
	"github.com/asgari-hamid/schemagen/payloads"
	"github.com/dave/jennifer/jen"
)

func main() {
	//generatePayload()
	usePayload()
}

func usePayload() {
	fields := []string{"id", "title", "price"}
	p := &payloads.Product{
		Mask:        payloads.BuildProductFieldMask(fields),
		Id:          "908ryfye89r7y",
		Title:       "TV",
		Description: "A TV",
		Price:       100,
		Weight:      21000,
		Published:   true,
	}
	bytes, _ := json.Marshal(p)
	fmt.Println(string(bytes))
}

func generatePayload() {
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
	gen.GeneratePayloadMarshaler(j, p)

	if err := j.Save("./payloads/product_gen.go"); err != nil {
		panic(err)
	}
}
