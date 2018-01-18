package model

import (
	"gopkg.in/mgo.v2/bson"
)

const(
	JD = "jd"
	TM = "tmall"
	TB = "taobao"
)
type Rule struct {
	Id bson.ObjectId `json:"id" bson:"_id"`
	Platform string  //平台
	Title string   //规则标题
	Url string   //规则地址
	Html string  //截取的html块的规则
	Name string  //爬商品名称的规则
	Sku string    //爬商品id的规则
	Price string   //爬商品梅林的规则
}

type RuleRepository interface {

	//保存规则(新增，更新)
	Save(rule Rule) (Rule, error)

	//根据id删除规则
	Delete(id bson.ObjectId) error

	//根据id更新规则
	Update(rule Rule) (Rule,error)

	//根据id查找规则
	Find(id bson.ObjectId) (Rule,error)

	//根据规则标题模糊搜索规则
	Search(title string) Rule
}
/*
  京东的规则
*/
func Jd() Rule{
	return Rule{Platform:JD,Html:"div[id='J_goodsList']",Name:"div[class='p-name p-name-type-2'] a em",Sku:"div[class='p-price'] strong",Price:"div[class='p-price'] strong i"}
}

/*
  天猫的规则
*/
func Tmall() Rule{
	return Rule{Platform:TM,Html:"div[id='J_ItemList']",Name:"div[class='productTitle productTitle-spu'] a",Sku:"",Price:"p[class='productPrice'] em"}
}

/*
  淘宝的规则
*/
func Tb() Rule{
	return Rule{Platform:TB,Html:"",Name:"",Sku:"",Price:""}
}