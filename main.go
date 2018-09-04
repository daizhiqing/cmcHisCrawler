package main

import (
	"cmcGoHis/cmc"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(3)

	//爬取CMC Week数据
	cmc.ScanCmcHistory()

	//爬取CMC币种日K入库
	list := cmc.ListCmcId()
	for _, name := range list {
		cmc.ScanDayHis(name, "20130425", time.Now().Format("20060102"))
	}
}
