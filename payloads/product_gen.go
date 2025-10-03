package payloads

import (
	"github.com/asgari-hamid/jsonw"
	"github.com/samber/mo"
)

const (
	ProductIdRef int = iota
	ProductTitleRef
	ProductDescriptionRef
	ProductPriceRef
	ProductWeightRef
	ProductPublishedRef
	ProductFieldCount
)

func BuildProductFieldMask(fields []string) []bool {
	mask := make([]bool, ProductFieldCount)
	for _, field := range fields {
		switch field {
		case "id":
			mask[ProductIdRef] = true
		case "title":
			mask[ProductTitleRef] = true
		case "description":
			mask[ProductDescriptionRef] = true
		case "price":
			mask[ProductPriceRef] = true
		case "weight":
			mask[ProductWeightRef] = true
		case "published":
			mask[ProductPublishedRef] = true
		}
	}
	return mask
}

type Product struct {
	Mask        []bool
	Id          string
	Title       string
	Description mo.Option[string]
	Price       float64
	Weight      mo.Option[int64]
	Published   bool
}

func (x *Product) writeJson(writer *jsonw.ObjectWriter, mask []bool) {
	writer.Open()
	if mask[ProductIdRef] {
		writer.StringField("id", x.Id)
	}
	if mask[ProductTitleRef] {
		writer.StringField("title", x.Title)
	}
	if mask[ProductDescriptionRef] {
		if value, exists := x.Description.Get(); exists {
			writer.StringField("description", value)
		} else {
			writer.NullField("description")
		}
	}
	if mask[ProductPriceRef] {
		writer.FloatField("price", x.Price)
	}
	if mask[ProductWeightRef] {
		if value, exists := x.Weight.Get(); exists {
			writer.IntegerField("weight", value)
		} else {
			writer.NullField("weight")
		}
	}
	if mask[ProductPublishedRef] {
		writer.BooleanField("published", x.Published)
	}
	writer.Close()
}
func (x *Product) MarshalJSON() ([]byte, error) {
	writer := jsonw.NewObjectWriter(nil)
	x.writeJson(writer, x.Mask)
	return writer.BuildBytes()
}
