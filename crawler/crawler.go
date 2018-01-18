package main

import(
	"depreciate/analyzer"
	"depreciate/model"
	"fmt"
)
func main(){

	//rule := model.Jd()
	//rule.Url = "https://search.jd.com/Search?keyword=手机&enc=utf-8"
	rule := model.Tmall()
	rule.Url = "https://list.tmall.com/search_product.htm?q=手机"
	goods := analyzer.Analyze(rule)
	fmt.Println(goods)
}