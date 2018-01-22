package analyzer

import (
	"github.com/gocolly/colly"
	"depreciate/model"
	"fmt"
	"strconv"
	"time"
	"depreciate/util"
	"github.com/PuerkitoBio/goquery"
)


func Analyze(rule model.Rule) []model.Goods{

	analyzer := colly.NewCollector()

	var goodss []model.Goods
	switch rule.Platform{
		case model.JD:
			JDAnalyze(analyzer,rule,&goodss)
		case model.TM:
			TMAnalyze(analyzer,rule,&goodss)
		default:
			fmt.Println("can not find analyzer")
	}

	analyzer.OnRequest(func(request *colly.Request) {
		fmt.Println("visiting", request.URL.String())
	})

	analyzer.OnError(func(response *colly.Response,err error) {
		fmt.Println("visit error happend...")
	})
	analyzer.Visit(rule.Url)
	return goodss
}


func JDAnalyze(analyzer *colly.Collector,rule model.Rule,goodss *[]model.Goods){

	analyzer.OnHTML(rule.Html, func(html *colly.HTMLElement) {
		html.DOM.Find("li.gl-item").Each(func(_ int,s *goquery.Selection){
			name := s.Find(rule.Name).Text()
			sku,_ := s.Attr("data-sku")
			priceStr := s.Find(rule.Price).Text()
			price,_ := strconv.ParseFloat(priceStr,2)
			time := time.Now().Format(util.TimeLayout)
			prices := []model.Price{model.Price{Price:price,Time:time}}
			goods := model.Goods{Name:name,Sku:sku,Prices:prices}
			*goodss = append(*goodss,goods)
		})
	})
}

func TMAnalyze(analyzer *colly.Collector,rule model.Rule,goodss *[]model.Goods){

	analyzer.OnHTML(rule.Html, func(html *colly.HTMLElement) {
		html.DOM.Find("div.product").Each(func(_ int,s *goquery.Selection){
			name := s.Find(rule.Name).Text()
			sku,_ := s.Attr("data-id")
			priceStr,_ := s.Find(rule.Price).Attr("title")
			price,_ := strconv.ParseFloat(priceStr,2)
			time := time.Now().Format(util.TimeLayout)
			prices := []model.Price{model.Price{Price:price,Time:time}}
			goods := model.Goods{Name:name,Sku:sku,Prices:prices}
			*goodss = append(*goodss,goods)
		})
	})
}
