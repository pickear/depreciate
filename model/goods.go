package model

import (
	"gopkg.in/mgo.v2/bson"
)

type Goods struct {
	Id bson.ObjectId `json:"id" bson:"_id"`
	Name string
	Sku string
	Prices []Price
}

type Price struct{
	Price float64
	Time string
}



type GoodsRepository interface{

	//新增商品
	Insert(rule Goods) (Goods, error)

	//根据sku删除商品
	Delete(sku string) error

	//根据id查找商品
	Find(sku string) (Goods,error)

	//根据商品名模糊搜索商品
	Search(name string) []Goods
}