package main

import (
	"github.com/wspl/creeper"
	"github.com/facebookgo/inject"
	"depreciate/repository"
	_ "depreciate/user"
	"depreciate/rule"
	"fmt"
	"depreciate/goods"
	"strconv"
)


func main(){

	var graph inject.Graph

	userRepository := repository.NewUserRepository()
	ruleRepository := repository.NewRuleRepository()
	session := repository.MongoConnect()
	defer session.Close()
	err := graph.Provide(
		&inject.Object{Value: userRepository},
		&inject.Object{Value: ruleRepository},
		&inject.Object{Value: session},
	)

	if(err != nil){
		fmt.Println(err)
	}

	if err = graph.Populate();err != nil{
		fmt.Println(err)
    }

	r := rule.Jd()
	r.Title = "京东moto"
	r.Url = "https://search.jd.com/Search?keyword=moto%20z2%20play&enc=utf-8&page={@page}"
	
	rule := "page(@page=1) = \""+r.Url+"\"\n" +
	"news[]: page -> $(\"li.gl-item\")\n" +
	"    name: $(\""+r.Name+"\").text\n" +
	"    sku: $(\""+r.Sku+"\").class\n" +
	"    price: $(\""+r.Price+"\").text\n"

	c := creeper.New(rule)
	c.Array("news").Each(func(c *creeper.Creeper) {
		price,err := strconv.ParseFloat(c.String("price"),2)
		if err != nil{
			fmt.Println("价格转换失败")
			return
		}
		goods := goods.Goods{Name:c.String("name"),Sku:c.String("sku"),Price:price}

		fmt.Println(goods)
		fmt.Println("===")
	})
}