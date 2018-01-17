package modle

import (
	"gopkg.in/mgo.v2/bson"
)

type Goods struct {
	Id bson.ObjectId `json:"id" bson:"_id"`
	Name string
	Sku string
	Price float64
}