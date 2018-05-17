package main

import(
	"depreciate/analyzer"
	"depreciate/model"
	"fmt"
)
func main(){

	//url := "https://search.jd.com/Search?keyword=手机&enc=utf-8"
	//url := "https://list.tmall.com/search_product.htm?q=手机"
    url := "http://www.gzruyue.org.cn:8094/api/Product/ProductDayArrayList?pid=5124546980941518626"
	goods := analyzer.Analyze(model.RY,url)
	fmt.Println(goods)
}