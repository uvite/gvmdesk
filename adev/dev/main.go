package main

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
	"github.com/uvite/gvmbot/pkg/cmd/cmdutil"
	gtools "github.com/uvite/gvmdesk/controller"
	"github.com/uvite/gvmdesk/internal"
	"io/ioutil"
	"net/http"
	"strings"
	"syscall"
)

func getBinanceStreams() {
	//streams = "/stream?streams="
	url := "https://api.binance.com/api/v3/ticker/price"
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	fmt.Println(err)
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	symbols := gjson.Get(string(body), "#.symbol").Array()

	ss := []string{}
	for _, symbol := range symbols {

		if symbol.String()[len(symbol.String())-4:] == "USDT" && !strings.Contains(symbol.String(), "UP") && !strings.Contains(symbol.String(), "DOWN") {
			//streams = streams + strings.ToLower(symbol.String()) + "@kline_1m" + "/"
			ss = append(ss, symbol.String())
			fmt.Println(symbol)
		}

	}
	fmt.Println(strings.Join(ss, ","))

	//streams = streams[:len(streams)-1]
	//streams = streams[:len(streams)-1]
	return
}
func main() {

	//getBinanceStreams()
	app := gtools.NewApp()
	ctx := context.Background()
	app.Ctx = ctx
	app.InitExchange()
	symbols := []string{"BTCUSDT", "ETHUSDT", "BNBUSDT", "BCCUSDT", "NEOUSDT", "LTCUSDT", "QTUMUSDT", "ADAUSDT", "XRPUSDT", "EOSUSDT"}
	//	"TUSDUSDT", "IOTAUSDT", "XLMUSDT", "ONTUSDT", "TRXUSDT", "ETCUSDT", "ICXUSDT", "VENUSDT", "NULSUSDT", "VETUSDT"}
	intervals := []string{"1m", "3m", "5m", "15m", "30m", "1h", "4h", "8h", "1d"}

	for _, symbol := range symbols {
		for _, interval := range intervals {
			app.AddSymbolInterval(symbol, interval)

			item := internal.AlertItem{
				Title:    "rsi",
				Symbol:   symbol,
				Interval: interval,
				Path:     "/hein/gvmdesk/js/4.js",
			}
			 app.RunTestFile(item)
		}
		//fmt.Println(k)
		//app.AddSymbolInterval(k, "1m")
		//app.AddSymbolInterval(k, "3m")
		//app.AddSymbolInterval(k, "5m")
		//app.AddSymbolInterval(k, "15m")
	}

	if err := app.Exchange.Stream.Connect(app.Ctx); err != nil {
		fmt.Println(err)
	}

	log.Infof("connecting...")

	cmdutil.WaitForSignal(ctx, syscall.SIGINT, syscall.SIGTERM)

}
