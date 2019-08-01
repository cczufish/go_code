package main

import (
	"fmt"
	"net/http"
	"github.com/PuerkitoBio/goquery"
)

func main() {
	res, err := http.Get("http://www.dzkbw.com/go.asp?hide=true&t=1&url=iuuq;00xxx/zvxfo{jzvbo/dpn0skc02t0e{lc033543/iunm")
	defer res.Body.Close()
	if err != nil {
		fmt.Println("error1", err)
	}
	utfBody, err := iconv.NewReader(res.Body, "gbk", "utf-8")
	if err != nil {
		fmt.Println("error2", err)
	}

	doc, err := goquery.NewDocumentFromReader(utfBody)
	if err != nil {
		fmt.Println("error3", err)
	}
	fmt.Println(doc)
}