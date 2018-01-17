package repository

import(
	"depreciate/modle"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)
type RuleRepositoryImpl struct{
	Mgo *mgo.Session `inject:""`
}

func (repository RuleRepositoryImpl) Save(rule modle.Rule) (modle.Rule, error){

	if len(rule.Id) <= 0{
		rule.Id = bson.NewObjectId()
	}
	fmt.Println("rule save")
	_,err := repository.collection().Upsert(bson.M{"_id":rule.Id},rule)
	if(err != nil){
		fmt.Println(err)
		return rule,err
	}
	return rule,nil
}

func (repository RuleRepositoryImpl) Delete(id bson.ObjectId) error{
	fmt.Println("user rule")
	err := repository.collection().Remove(bson.M{"_id":id})
	if err != nil{
		fmt.Println(err)
		return err
	}
	return nil
}

func (repository RuleRepositoryImpl) Update(rule modle.Rule) (modle.Rule,error){
	fmt.Println("user rule")

	err := repository.collection().Update(bson.M{"_id":rule.Id},rule)
	if err != nil{
		fmt.Println(err)
		return rule,err
	}
	return rule,nil
}

func (repository RuleRepositoryImpl) Find(id bson.ObjectId) modle.Rule{
	var rule modle.Rule
	err := repository.collection().Find(bson.M{"_id":id}).One(&rule)
	if(err != nil){
		fmt.Println(err)
	}
	return rule
}

func (repository RuleRepositoryImpl) Search(title string) modle.Rule{
	var rule modle.Rule
	err := repository.collection().Find(bson.M{"title":title}).One(&rule)
	if(err != nil){
		fmt.Println(err)
	}
	return rule
}

func (ur RuleRepositoryImpl) collection() *mgo.Collection{
	collection := ur.Mgo.DB("depreciate").C("rule")
	return collection
}

func NewRuleRepository() *RuleRepositoryImpl{

	return &RuleRepositoryImpl{}
}