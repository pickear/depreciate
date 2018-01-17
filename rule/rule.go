package rule

import (
	"gopkg.in/mgo.v2/bson"
)

type Rule struct {
	Id bson.ObjectId `json:"id" bson:"_id"`
	Title string   //规则标题
	Url string   //规则地址
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
	Find(id bson.ObjectId) Rule

	//根据规则标题模糊搜索规则
	Search(title string) Rule
}
/*
  京东的规则
*/
func Jd() *Rule{
	return &Rule{Name:"div.p-name a em",Sku:"div.p-price strong",Price:"div.p-price strong i"}
}

/*
  天猫的规则
*/
func Tmall() *Rule{
	return &Rule{Name:"",Sku:"",Price:""}
}

/*
  淘宝的规则
*/
func Tb() *Rule{
	return &Rule{Name:"",Sku:"",Price:""}
}