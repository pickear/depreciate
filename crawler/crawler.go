package main

import(
	"depreciate/analyzer"
	"depreciate/model"
	"fmt"
	"strconv"
	"bytes"
	"time"
	"net/smtp"
	"strings"
)
func main(){

	//url := "https://search.jd.com/Search?keyword=手机&enc=utf-8"
	//url := "https://list.tmall.com/search_product.htm?q=手机"
	ticker := time.NewTicker(time.Minute * 15)
	for _ = range ticker.C{
		doAnalyze()
	}
}

func doAnalyze()  {
	url := "http://www.gzruyue.org.cn:8094/api/Product/ProductDayArrayList?pid=5124546980941518626"
	ruyueReponse := analyzer.Analyze(model.RY,url)
	var message bytes.Buffer
	tellFollwer := false
	_ruyueReponse := ruyueReponse.(*model.RuYueResponse)
	message.WriteString("班次: "+_ruyueReponse.Data.Product.Pnm+"\r\n")
	message.WriteString("经过站点: "+_ruyueReponse.Data.Product.Rdc+"\r\n")

	items := _ruyueReponse.Data.Items
	for _,item := range items{
		message.WriteString(item.Date+"\r\n")
		clsinfs := item.Clsinf
		for _,clsinf := range clsinfs{

			message.WriteString("\t"+clsinf.Clstm+" : ")
			seats,_ := strconv.Atoi(clsinf.Seats)
			if seats > 0{
				dateTime,_ := time.ParseInLocation("2006-01-02",item.Date,time.Local)
				if dateTime.After(time.Now()){
					tellFollwer = true
				}
				message.WriteString("有"+strconv.Itoa(seats)+"张票\r\n")
			}else {
				message.WriteString("没票"+"\r\n")
			}
		}
	}

	if tellFollwer{
		fmt.Println("将会发消息通知")
		host := "smtp.qq.com:25"
		user := "114231159@qq.com"
		password := "aqixxeeslpfqcajb"
		to := "114231159@qq.com"
		send_to := strings.Split(to, ";")
		content_type := "Content-Type: text/plain; charset=UTF-8"
		subject := "如约购票通知"
		auth := smtp.PlainAuth("",user,password,"smtp.qq.com")
		msg := []byte("To: " + strings.Join(send_to, ",") + "\r\nFrom: " + user +
			"<" + user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + message.String())
		error := smtp.SendMail(host,auth,user,send_to, msg)
		if nil != error{
			fmt.Println(error)
		}
	}
	fmt.Println(message.String())
}