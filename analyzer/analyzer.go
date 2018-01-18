package analyzer

import (
	"github.com/gocolly/colly"
	"depreciate/model"
	"fmt"
	"strconv"
	"time"
	"depreciate/util"
)


func Analyze(rule model.Rule) model.Goods{

	analyzer := colly.NewCollector()

	var goods model.Goods
	analyzer.OnHTML(rule.Name, func(e *colly.HTMLElement) {
		goods.Name = e.Text
	})
	analyzer.OnHTML(rule.Sku, func(e *colly.HTMLElement) {
		goods.Sku = e.Attr("class")
	})
	analyzer.OnHTML(rule.Price, func(e *colly.HTMLElement) {
		price,_ := strconv.ParseFloat(e.Text,2)
		time := time.Now().Format(util.TimeLayout)
		goods.Prices = []model.Price{model.Price{Price:price,Time:time}}
	})
	analyzer.OnRequest(func(request *colly.Request) {
		fmt.Println("visiting", request.URL.String())
	})

	analyzer.Visit(rule.Url)
	return goods
}