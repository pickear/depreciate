package main

import (
	"github.com/wspl/creeper"
	"github.com/facebookgo/inject"
	"depreciate/repository"
	"depreciate/modle"
	"fmt"
	"strconv"
	"time"
	"depreciate/util"
)


func main(){

	var graph inject.Graph

	userRepository := repository.NewUserRepository()
	ruleRepository := repository.NewRuleRepository()
	goodsRepository := repository.NewGoodsRepository()
	session := repository.MongoConnect()
	defer session.Close()
	err := graph.Provide(
		&inject.Object{Value: userRepository},
		&inject.Object{Value: ruleRepository},
		&inject.Object{Value: goodsRepository},
		&inject.Object{Value: session},
	)

	if(err != nil){
		fmt.Println(err)
	}

	if err = graph.Populate();err != nil{
		fmt.Println(err)
    }

	r := modle.Jd()
	r.Title = "京东moto"
	r.Url = "https://search.jd.com/Search?keyword=手机&enc=utf-8&page={@page}"
	
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
		time := time.Now().Format(util.TimeLayout)
		prices := []modle.Price{modle.Price{Price:price,Time:time}}
		goods := modle.Goods{Name:c.String("name"),Sku:c.String("sku"),Prices:prices}
		goodsRepository.Save(goods)
		fmt.Println(goods)
		fmt.Println("===")
	})
}