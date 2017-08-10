package main

import (
	"fmt"
	"github.com/anaskhan96/soup"
	"os"
	//"github.com/djimenez/iconv-go"
	//"github.com/PuerkitoBio/goquery"
	//"strconv"
	//"go/doc"
)

func main() {
	resp, err := soup.Get("http://news.baidu.com/n?cmd=4&class=civilnews&pn=1")
	if err != nil {
		os.Exit(1)
	}

	//utfBody, err := iconv.NewReader(resp.Body, "gb2312", "utf-8")
	//if err != nil {
	//	fmt.Println(err.Error())
	//}

	//doc := soup.HTMLParse(resp)
	doc := soup.HTMLParseGBK(resp)

	//links := doc.Find("div", "id", "comicLinks").FindAll("a")
	links := doc.Find("div", "class", "p2").FindAll("a")
	for _, link := range links {
		fmt.Println("Title :", link.Text())
		fmt.Println("Link :", link.Attrs()["href"])

		span := doc.Find("div", "class", "p2").Find("span")
		fmt.Println("Span :", span.Text())

	}

	//spans := doc.Find("div", "class", "p2").FindAll("span")
	//for _, span := range spans {
	//	fmt.Println("Span :", span.Text())
	//}
}
