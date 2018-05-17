package analyzer

import (
	"depreciate/model"
	"github.com/gocolly/colly"
	"fmt"
	"encoding/json"
)

type RuYueAnalyzer struct {
	collector *colly.Collector
}


func (ruYueAnalyzer RuYueAnalyzer) analyze(rule model.Rule)(interface{})  {
	mapResult := make(map[string]interface{})
	shifts := model.Shifts{}
	ruYueAnalyzer.collector.OnResponse(func(response *colly.Response) {
		err := json.Unmarshal(response.Body[:],&mapResult)
		if err != nil{
			println(err)
		}

		data,_ := mapResult["data"]
		product,_ := data.(map[string]interface{})["Product"]
		shifts.Pnm,_ = string.(product.(map[string]interface{})["pnm"])

		fmt.Printf("%s\n", mapResult)
	})
	return &shifts
}
