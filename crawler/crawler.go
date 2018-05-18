package main

import(
	"depreciate/analyzer"
	"depreciate/model"
	"fmt"
	"strconv"
	"bytes"
	"time"
)
func main(){

	//url := "https://search.jd.com/Search?keyword=手机&enc=utf-8"
	//url := "https://list.tmall.com/search_product.htm?q=手机"
    url := "http://www.gzruyue.org.cn:8094/api/Product/ProductDayArrayList?pid=5124546980941518626"
	ruyueReponse := analyzer.Analyze(model.RY,url)
	var buf bytes.Buffer
	tellFollwer := false
	_ruyueReponse := ruyueReponse.(*model.RuYueResponse)
	buf.WriteString("班次: "+_ruyueReponse.Data.Product.Pnm+"\r\n")
	buf.WriteString("经过站点: "+_ruyueReponse.Data.Product.Rdc+"\r\n")

	items := _ruyueReponse.Data.Items
	for _,item := range items{
		buf.WriteString(item.Date+"\r\n")
		clsinfs := item.Clsinf
		for _,clsinf := range clsinfs{

			buf.WriteString("\t"+clsinf.Clstm+" : ")
			seats,_ := strconv.Atoi(clsinf.Seats)
			if seats > 0{
				dateTime,_ := time.ParseInLocation("2006-01-02",item.Date,time.Local)
				if dateTime.After(time.Now()){
					tellFollwer = true
				}
				buf.WriteString("有"+strconv.Itoa(seats)+"张票\r\n")
			}else {
				buf.WriteString("没票"+"\r\n")
			}
		}
	}

	if tellFollwer{
		fmt.Println("将会发消息通知")
	}
	fmt.Println(buf.String())
}