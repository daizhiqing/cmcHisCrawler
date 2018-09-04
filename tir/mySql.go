package tir

import (
	"cmcGoHis/model"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	db, _ = sql.Open("mysql", "user:pwd@tcp(host:port)/db?charset=utf8&loc=Asia%2FShanghai&parseTime=true")
	db.SetMaxOpenConns(1000)
	db.SetMaxIdleConns(10)
	db.Ping()
}

type CurrencyInfoHBase struct {
	id         int
	rowKey     string
	cmcId      string
	createTime time.Time
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

//插入货币周价格数据
func InsertCmcHisWeek(cmc model.CmcHisWeekData) {
	stmt, err := db.Prepare("INSERT cmc_his_week_usd (date, sortId, cmcId, symbol, marketCap, price, circulatingSupply, volume24, h1Change, h24Change, d7Change) values (?,?,?,?,?,?,?,?,?,?,?)")
	defer stmt.Close()
	CheckErr(err)

	res, err := stmt.Exec(
		cmc.Date,
		cmc.SortId,
		cmc.CmcId,
		cmc.Symbol,
		cmc.MarketCap,
		cmc.Price,
		cmc.CirculatingSupply,
		cmc.Volume24,
		cmc.H1Change,
		cmc.H24Change,
		cmc.D7Change)
	CheckErr(err)

	id, err := res.LastInsertId()
	CheckErr(err)
	fmt.Println("cmc历史周数据插入完成", id)
}

//插入货币日K价格数据
func InsertCmcHisDayKline(cmc model.CmcDayHisKline) {
	stmt, err := db.Prepare("INSERT cmc_his_day_kline_usd (cmcId , date, openPrice, highPrice, lowPrice, closePrice, volume, marketCap) values (?,?,?,?,?,?,?,?)")
	defer stmt.Close()
	CheckErr(err)
	res, err := stmt.Exec(
		cmc.CmcId,
		cmc.Date,
		cmc.OpenPrice,
		cmc.HighPrice,
		cmc.LowPrice,
		cmc.ClosePrice,
		cmc.Volume,
		cmc.MarketCap)
	id, err := res.LastInsertId()
	CheckErr(err)
	fmt.Println("cmc历史日K数据插入完成", id)
}
