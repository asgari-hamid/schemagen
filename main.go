package main

import (
	"encoding/json"
	"fmt"

	"github.com/asgari-hamid/schemagen/code"
	"github.com/asgari-hamid/schemagen/gen"
	"github.com/asgari-hamid/schemagen/payloads"
	"github.com/dave/jennifer/jen"
	"github.com/samber/mo"
)

func main() {
	//generateProductPayload()
	//generateProductResponsePayload()
	usePayload()
}

func usePayload() {
	fields := []string{"id", "title", "description", "price", "weight", "published"}
	p := payloads.Product{
		Mask:        payloads.BuildProductFieldMask(fields),
		Id:          "908ryfye89r7y",
		Title:       "TV",
		Description: mo.Some("A TV"),
		Price:       100,
		Weight:      mo.None[int64](),
		Published:   true,
	}
	r := &payloads.ProductResponse{
		Product: mo.Some(p),
	}
	bytes, _ := json.Marshal(r)
	fmt.Println(string(bytes))
}

func generateProductPayload() {
	j := jen.NewFile("payloads")

	product := &code.Payload{
		Name: "Product",
		Fields: []*code.PayloadField{
			{StructName: "Id", JsonName: "id", Type: code.SchemaTypeString, Tags: nil},
			{StructName: "Title", JsonName: "title", Type: code.SchemaTypeString, Tags: nil},
			{StructName: "Description", JsonName: "description", Type: code.SchemaTypeString, Nullable: true, Tags: nil},
			{StructName: "Price", JsonName: "price", Type: code.SchemaTypeNumber, Tags: nil},
			{StructName: "Weight", JsonName: "weight", Type: code.SchemaTypeInteger, Nullable: true, Tags: nil},
			{StructName: "Published", JsonName: "published", Type: code.SchemaTypeBoolean, Tags: nil},
		},
	}

	gen.GenerateImports(j)
	gen.GeneratePayloadFieldIndices(j, product)
	gen.GeneratePayloadFieldMask(j, product)
	gen.GeneratePayloadStruct(j, product)
	gen.GeneratePayloadJsonWriter(j, product)
	gen.GeneratePayloadMarshaler(j, product)

	if err := j.Save("./payloads/product_gen.go"); err != nil {
		panic(err)
	}
}

func generateProductResponsePayload() {
	j := jen.NewFile("payloads")

	productResponse := &code.Payload{
		Name: "ProductResponse",
		Fields: []*code.PayloadField{
			{StructName: "Product", JsonName: "product", Type: code.SchemaTypeObject, TypeName: "Product", Nullable: true, Tags: nil},
		},
	}

	gen.GenerateImports(j)
	gen.GeneratePayloadFieldIndices(j, productResponse)
	gen.GeneratePayloadFieldMask(j, productResponse)
	gen.GeneratePayloadStruct(j, productResponse)
	gen.GeneratePayloadJsonWriter(j, productResponse)
	gen.GeneratePayloadMarshaler(j, productResponse)

	if err := j.Save("./payloads/product_response_gen.go"); err != nil {
		panic(err)
	}
}
