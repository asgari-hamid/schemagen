package main

import (
	"schemagen/code"
	"schemagen/gen"

	"github.com/dave/jennifer/jen"
)

func main() {
	j := jen.NewFile("payloads")

	p := &code.Payload{
		Name: "Product",
		Fields: []*code.PayloadField{
			{Name: "Id", Type: "string", Tags: nil},
			{Name: "Title", Type: "string", Tags: nil},
			{Name: "Description", Type: "string", Tags: nil},
			{Name: "Price", Type: "float64", Tags: nil},
		},
	}

	gen.GeneratePayloadStruct(j, p)

	if err := j.Save("./payloads/product_gen.go"); err != nil {
		panic(err)
	}
}
