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

type Analyzer interface {
	analyze(rule model.Rule) interface{}
}

func Analyze(platform,url string) interface{}{

	collector := colly.NewCollector()
	var analyzer Analyzer
	var rule model.Rule
	switch platform {
	case model.JD:
		analyzer = JDAnalyzer{collector}
		rule = model.Jd(url)
	case model.TM:
		analyzer = TMAnalyzer{collector}
		rule = model.Tmall(url)
	case model.RY:
		analyzer = RuYueAnalyzer{collector}
		rule = model.RuYue(url)
	default:
		fmt.Println("can not find analyzer")
	}

	result := analyzer.analyze(rule)

	collector.OnRequest(func(request *colly.Request) {
		fmt.Println("visiting", request.URL.String())
	})

	collector.OnError(func(response *colly.Response,err error) {
		fmt.Println("visit error happend...")
	})
	collector.Visit(rule.Url)
	return result
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
