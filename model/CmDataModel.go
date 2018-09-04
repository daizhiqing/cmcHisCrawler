package model

//周数据历史实体
type CmcHisWeekData struct {
	Date              string
	SortId            string
	CmcId             string
	Symbol            string
	MarketCap         string
	Price             string
	CirculatingSupply string
	Volume24          string
	H1Change          string
	H24Change         string
	D7Change          string
}

//币种的信息
type CmcEntity struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Symbol       string `json:"symbol"`
	Website_slug string `json:"website_slug"`
}

type CmcListingData struct {
	Data     []CmcEntity `json:"data"`
	Metadata interface{} `json:"metadata"`
}

//日K数据
type CmcDayHisKline struct {
	CmcId      string
	Date       string
	OpenPrice  string
	HighPrice  string
	LowPrice   string
	ClosePrice string
	Volume     string
	MarketCap  string
}
