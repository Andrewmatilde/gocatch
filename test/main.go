package main

import (
	"github.com/PuerkitoBio/goquery"
	"gocatch"
	"net/url"
	"os"
	"time"
)

type DLoader struct {}

func (d *DLoader) Work(
	u *gocatch.StrStack) gocatch.ResPipe {
	s := u.SafePopO()
	resp,_ := goquery.NewDocument(s)
	ER := gocatch.ResPipe{Res: resp,Data:s}
	return ER
}

type Spider struct {}

func (sp *Spider)Analyze(s *gocatch.StrStack, er gocatch.ResPipe) {
	doc := er.GetResValueInterface().(*goquery.Document)
	ResUrl := er.Data
	ParsedUrl, _ := url.Parse(ResUrl)
	row,_ := url.ParseQuery(ParsedUrl.RawQuery)
	data := doc.Find("div.historyList>table:first-child>tbody>tr")
	file, _ := os.OpenFile(row["hy"][0]+row["page"][0]+".txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	defer file.Close()
	data.Each(func(i int, selection *goquery.Selection) {
		if i>0 {
			selection.Find("td").Each(func(i int, selection *goquery.Selection){
				_, _ = file.WriteString(selection.Text() + "," + row["hy"][0])
			})
			_, _ = file.WriteString("\n")
		}
	})
}

func main() {
	s := gocatch.StrStack{}
	urls := []string{"http://vip.stock.finance.sina.com.cn/q/view/vFutures_History.php?page=1&breed=CU0&start=2010-01-01&end=2020-01-01&jys=shfe&pz=CU&hy=CU0&type=inner&name=&#161;&#228;%C2%A1%C2%A7&#174;",
		"http://vip.stock.finance.sina.com.cn/q/view/vFutures_History.php?page=2&breed=CU0&start=2010-01-01&end=2020-01-01&jys=shfe&pz=CU&hy=CU0&type=inner"}
	s.SafePushA(urls)
	dl := DLoader{}
	var dm gocatch.WorkMachine = &dl
	sp := Spider{}
	var sa gocatch.AnalyzeMachine = &sp
	m := gocatch.Manager{WorkLineNum: 2, AnaLineNum: 1,Worker:dm,Analyst:sa,Stack:&s}
	Wl := gocatch.CreatWorkLineList(m,time.Duration(100)*time.Millisecond,false,false)
	Al := gocatch.CreateAnaLineList(m,time.Duration(100)*time.Millisecond,false,false)
	gocatch.RunEngine(m,Wl,Al)
}