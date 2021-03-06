package types

import (
	"time"

	"github.com/ethereum/go-ethereum/common"
	validation "github.com/go-ozzo/ozzo-validation"
	"labix.org/v2/mgo/bson"
)

type Token struct {
	ID              bson.ObjectId `json:"id" bson:"_id"`
	Name            string        `json:"name" bson:"name"`
	Symbol          string        `json:"symbol" bson:"symbol"`
	Image           Image         `json:"image" bson:"image"`
	ContractAddress string        `json:"contractAddress" bson:"contractAddress"`
	Decimal         int           `json:"decimal" bson:"decimal"`
	Active          bool          `json:"active" bson:"active"`

	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
}

type Image struct {
	URL  string                 `json:"url" bson:"url"`
	Meta map[string]interface{} `json:"meta" bson:"meta"`
}

func (t Token) Validate() error {
	return validation.ValidateStruct(&t,
		validation.Field(&t.Name, validation.Required),
		validation.Field(&t.Symbol, validation.Required),
		validation.Field(&t.ContractAddress, validation.Required, validation.NewStringRule(common.IsHexAddress, "Invalid Address")),
		validation.Field(&t.Decimal, validation.Required),
	)
}
