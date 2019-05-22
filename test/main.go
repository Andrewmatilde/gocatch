package main

import (
	"github.com/PuerkitoBio/goquery"
	"gocatch"
	"net/url"
	"os"
)

type DLoad struct {}

func (d *DLoad) Down(
	u *gocatch.Stack) gocatch.EleRes {
	s := u.SafePopO()
	resp,_ := goquery.NewDocument(s)
	ER := gocatch.EleRes{Res:resp,Data:s}
	return ER
}

type Spide struct {}

func (sp *Spide)Analyze(s *gocatch.Stack, er gocatch.EleRes) {
	doc := er.GetResValueNeedChangeTypes().(*goquery.Document)
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
	s := gocatch.Stack{}
	urls := []string{"http://vip.stock.finance.sina.com.cn/q/view/vFutures_History.php?page=1&breed=CU0&start=2010-01-01&end=2020-01-01&jys=shfe&pz=CU&hy=CU0&type=inner&name=&#161;&#228;%C2%A1%C2%A7&#174;",
		"http://vip.stock.finance.sina.com.cn/q/view/vFutures_History.php?page=2&breed=CU0&start=2010-01-01&end=2020-01-01&jys=shfe&pz=CU&hy=CU0&type=inner"}
	s.SafePushA(urls)
	e := gocatch.Manager{DlLineNum: 2, AnaLineNum: 1}
	dl := DLoad{}
	var dm gocatch.DownMachine = &dl
	sp := Spide{}
	var as gocatch.Analysis = &sp
	e.RunEngine(&s,dm,as)
}