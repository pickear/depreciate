package main

import(
	"depreciate/analyzer"
	"depreciate/model"
	"fmt"
)
func main(){

	rule := model.Jd()
	rule.Url = "https://search.jd.com/Search?keyword=手机&enc=utf-8"
	goods := analyzer.Analyze(rule)
	fmt.Println(goods)
}