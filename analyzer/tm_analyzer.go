package analyzer

import (
	"depreciate/model"
	"github.com/gocolly/colly"
	"github.com/PuerkitoBio/goquery"
	"strconv"
	"time"
	"depreciate/util"
)

type TMAnalyzer struct {
	collector *colly.Collector
}

func (tmAnalyzer TMAnalyzer) analyze(rule model.Rule) (interface{}){

	var goodss []model.Goods
	tmAnalyzer.collector.OnHTML(rule.Html, func(html *colly.HTMLElement) {
		html.DOM.Find("div.product").Each(func(_ int,s *goquery.Selection){
			name := s.Find(rule.Name).Text()
			sku,_ := s.Attr("data-id")
			priceStr,_ := s.Find(rule.Price).Attr("title")
			price,_ := strconv.ParseFloat(priceStr,2)
			time := time.Now().Format(util.TimeLayout)
			prices := []model.Price{model.Price{Price:price,Time:time}}
			goods := model.Goods{Name:name,Sku:sku,Prices:prices}
			goodss = append(goodss,goods)
		})
	})
	return &goodss
}
