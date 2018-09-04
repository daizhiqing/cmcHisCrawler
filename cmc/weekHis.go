package cmc

import (
	"cmcGoHis/model"
	"cmcGoHis/tir"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const DOMAIN = "https://coinmarketcap.com"

var limit = make(chan int, 10)

//扫描CMC的历史数据
func ScanCmcHistory() {
	respone, err := http.Get(DOMAIN + "/historical/")

	if err != nil {
		log.Println(err.Error())
		return
	}

	defer respone.Body.Close()

	if respone.StatusCode == 200 {
		doc, err := goquery.NewDocumentFromReader(respone.Body)
		if err != nil {
			log.Fatal(err)
		}

		doc.Find(".container .row .col-lg-10 .row .col-sm-4").Each(func(i int, selection *goquery.Selection) {
			//查找所有链接
			hrefA := selection.Find("ul li a")

			hrefA.Each(func(i int, selection *goquery.Selection) {
				hrefText := selection.Text()
				hrefLink, ok := selection.Attr("href")
				if ok {
					limit <- 1
					fmt.Println("开始爬取：", hrefText, hrefLink)
					go CrawTableDate(hrefLink)
				}
			})

		})
	}
}

// 用于爬取二级页面数据
func CrawTableDate(url string) {
	urlSp := strings.Split(url, "/")

	//i , _ := strconv.Atoi(urlSp[2])
	//if i < 20180701 {
	//	<-limit
	//	return
	//}

	resp, err := http.Get(DOMAIN + url)
	//resp, err := http.Get("https://coinmarketcap.com/zh/historical/20180902/")

	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	doc.Find("table tbody tr").Each(func(i int, selection *goquery.Selection) {

		ch := make(chan string, 10)
		defer close(ch)
		selection.Find("td").Each(func(i int, selection *goquery.Selection) {
			btn := selection.Has("button")
			if len(btn.Nodes) == 0 {
				data, ok := selection.Attr("data-sort")
				if ok {
					ch <- strings.Replace(data, "\n", "", -1)
					//fmt.Println(data)
				} else {
					ch <- strings.Replace(selection.Text(), "\n", "", -1)
					//fmt.Println(selection.Text())
				}
			}
		})
		id, cmcId, symbol, marketCap, price, circulatingSupply, volume24, h1Change, h24Change, d7Change := <-ch, <-ch, <-ch, <-ch, <-ch, <-ch, <-ch, <-ch, <-ch, <-ch

		cmc := model.CmcHisWeekData{
			urlSp[2], id, cmcId, symbol, marketCap, price, circulatingSupply, volume24, h1Change, h24Change, d7Change,
		}
		fmt.Println(urlSp[2], id, cmcId, symbol, marketCap, price, circulatingSupply, volume24, h1Change, h24Change, d7Change)
		tir.InsertCmcHisWeek(cmc)
	})
	<-limit
}
