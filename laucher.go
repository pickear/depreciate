package main

import (
	"github.com/wspl/creeper"
)


func main(){

	rule := "page(@page=1) = \"https://search.jd.com/Search?keyword=摩托罗拉 Moto M(XT1662) 4G+32G 耀世金移动联通电信4G手机 双卡双待&enc=utf-8&page={@page}\"\n" +
	"news[]: page -> $(\"li.gl-item\")\n" +
	"    title: $(\"div.p-name a em\").text\n" +
	"    sku: $(\"div.p-price strong\").class\n" +
	"    price: $(\"div.p-price strong i\").text\n" +
	"    tips: $(\"div.p-commit a\").text"

	c := creeper.New(rule)
	c.Array("news").Each(func(c *creeper.Creeper) {
		println("商品名: ", c.String("title"))
		println("商品编号: ", c.String("sku"))
		println("商品价格: ", c.String("price"))
		println("评论数: ", c.String("tips"))
		println("===")
	})
}