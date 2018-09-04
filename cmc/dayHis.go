package cmc

import (
	"cmcGoHis/model"
	"cmcGoHis/tir"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

//获取CMC网站上所有的币种名称
func ListCmcId() []string {
	url := "https://api.coinmarketcap.com/v2/listings/"

	resp, err := http.Get(url)
	CheckErr(err)

	defer resp.Body.Close()
	var cmcIdList []string
	if resp.StatusCode == 200 {
		result, err := ioutil.ReadAll(resp.Body)
		CheckErr(err)
		respJsonObj := model.CmcListingData{}
		err = json.Unmarshal(result, &respJsonObj)
		CheckErr(err)
		fmt.Println(respJsonObj.Data)

		for _, data := range respJsonObj.Data {
			fmt.Println(data)
			cmcIdList = append(cmcIdList, data.Website_slug)
		}
	}
	return cmcIdList
}

//爬取CMC的日K数据
func ScanDayHis(cmcId, start, end string) {
	url := DOMAIN + "/zh/currencies/" + cmcId + "/historical-data/?start=" + start + "&end=" + end
	fmt.Println(url)
	respone, err := http.Get(url)

	CheckErr(err)
	defer respone.Body.Close()

	if respone.StatusCode == 200 {
		doc, err := goquery.NewDocumentFromReader(respone.Body)
		CheckErr(err)

		doc.Find(".table-responsive table tbody tr").Each(func(i int, selection *goquery.Selection) {

			ch := make(chan string, 7)
			defer close(ch)
			selection.Find("td").Each(func(i int, selection *goquery.Selection) {
				data, ok := selection.Attr("data-format-value")
				if ok {
					ch <- data
					//fmt.Println(data)
				} else {
					ch <- selection.Text()
					//fmt.Println(selection.Text())
				}
			})

			date, openPrice, highPrice, lowPrice, closePrice, volume, marketCap := <-ch, <-ch, <-ch, <-ch, <-ch, <-ch, <-ch

			//fmt.Println(date ,openPrice , highPrice , lowPrice , closePrice ,volume , marketCap)
			r := strings.NewReplacer("年", "", "月", "", "日", "")
			date = r.Replace(date)

			data := model.CmcDayHisKline{cmcId, date, openPrice, highPrice, lowPrice, closePrice, volume, marketCap}

			if "-" == marketCap {
				marketCap = "0"
			} else {
				// 格式化市值的科学计数法
				var temp float64
				_, err := fmt.Sscanf(marketCap, "%e", &temp)
				CheckErr(err)
				marketCap = fmt.Sprintf("%f", temp)
			}
			data.MarketCap = marketCap
			fmt.Println(data)
			tir.InsertCmcHisDayKline(data)

		})

	}
}
