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
	ruyueReponse := model.RuYueResponse{}
	ruYueAnalyzer.collector.OnResponse(func(response *colly.Response) {
		err := json.Unmarshal(response.Body[:],&ruyueReponse)
		if err != nil{
			println(err)
		}

		fmt.Printf("%s\n", ruyueReponse)
	})
	return &ruyueReponse
}
