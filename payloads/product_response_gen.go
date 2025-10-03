package payloads

import (
	"github.com/asgari-hamid/jsonw"
	"github.com/samber/mo"
)

const (
	ProductResponseProductRef int = iota
	ProductResponseFieldCount
)

func BuildProductResponseFieldMask(fields []string) []bool {
	mask := make([]bool, ProductResponseFieldCount)
	for _, field := range fields {
		switch field {
		case "product":
			mask[ProductResponseProductRef] = true
		}
	}
	return mask
}

type ProductResponse struct {
	Mask    []bool
	Product mo.Option[Product]
}

func (x *ProductResponse) WriteJson(writer *jsonw.ObjectWriter) {
	mask := x.Mask
	noMask := len(mask) != ProductResponseFieldCount
	writer.Open()
	if noMask || mask[ProductResponseProductRef] {
		if value, exists := x.Product.Get(); exists {
			obj := writer.ObjectField("product")
			value.WriteJson(obj)
		} else {
			writer.NullField("product")
		}
	}
	writer.Close()
}
func (x *ProductResponse) MarshalJSON() ([]byte, error) {
	writer := jsonw.NewObjectWriter(nil)
	x.WriteJson(writer)
	return writer.BuildBytes()
}
