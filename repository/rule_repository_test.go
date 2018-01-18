package repository

import (
	"testing"
	"depreciate/model"
	"github.com/facebookgo/inject"
	"fmt"
	"github.com/stretchr/testify/assert"
)

func TestRuleSave(t *testing.T){
	assert := assert.New(t)
	var graph inject.Graph

	ruleRepository := NewRuleRepository()
	session := MongoConnect()
	defer session.Close()
	err := graph.Provide(
		&inject.Object{Value: ruleRepository},
		&inject.Object{Value: session},
	)

	if(err != nil){
		fmt.Println(err)
	}

	if err = graph.Populate();err != nil{
		fmt.Println(err)
	}
	r := model.Rule{Title:"京东moto",Url:"https://search.jd.com/Search?keyword=摩托罗拉 Moto M(XT1662) 4G+32G 耀世金移动联通电信4G手机 双卡双待&enc=utf-8&page={@page}",Name:"div.p-name a em",Sku:"div.p-price strong",Price:"div.p-price strong i"}
	_,err = ruleRepository.Save(r)

	assert.Equal(err,nil,"save user err")
}


func TestRuleFind(t *testing.T){
	assert := assert.New(t)
	var graph inject.Graph

	ruleRepository := NewRuleRepository()
	session := MongoConnect()
	defer session.Close()
	err := graph.Provide(
		&inject.Object{Value: ruleRepository},
		&inject.Object{Value: session},
	)

	if(err != nil){
		fmt.Println(err)
	}

	if err = graph.Populate();err != nil{
		fmt.Println(err)
	}

	rule := ruleRepository.Search("京东moto")
	assert.Equal(rule.Title,"京东moto","search rule err")
}

func TestRuleUpdate(t *testing.T){
	assert := assert.New(t)
	var graph inject.Graph

	ruleRepository := NewRuleRepository()
	session := MongoConnect()
	defer session.Close()
	err := graph.Provide(
		&inject.Object{Value: ruleRepository},
		&inject.Object{Value: session},
	)

	if(err != nil){
		fmt.Println(err)
	}

	if err = graph.Populate();err != nil{
		fmt.Println(err)
	}

	rule := ruleRepository.Search("京东moto")
	rule.Title = "京东motorala"
	rule,err = ruleRepository.Update(rule)
	assert.Equal(err,nil,"update rule err")
}
