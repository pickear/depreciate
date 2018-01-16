package rule

import (
	"gopkg.in/mgo.v2/bson"
)

type Rule struct {
	Id bson.ObjectId `json:"id" bson:"_id"`
	Url string
	Name string
	Sku string
	Price string
}

/*
  京东的规则
*/
func (r Rule) Jd() *Rule{
	return &Rule{Name:"",Sku:"",Price:""}
}

/*
  天猫的规则
*/
func (r Rule) Tmall() *Rule{
	return &Rule{Name:"",Sku:"",Price:""}
}

/*
  淘宝的规则
*/
func (r Rule) Tb() *Rule{
	return &Rule{Name:"",Sku:"",Price:""}
}