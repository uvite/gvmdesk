package main

import (
	"context"
	"fmt"
	"github.com/adshao/go-binance/v2"
)

func main() {
	var (
		apiKey    = " "
		secretKey = " "
	)
	client := binance.NewClient(apiKey, secretKey)

	klines, err := client.NewKlinesService().Symbol("ETHUSDT").
		Interval("1m").Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, k := range klines {
		fmt.Println(k)
	}
	wsKlineHandler := func(event *binance.WsKlineEvent) {
		fmt.Println(event)
	}
	errHandler := func(err error) {
		fmt.Println(err)
	}
	doneC, _, err := binance.WsKlineServe("ETHUSDT", "1m", wsKlineHandler, errHandler)
	if err != nil {
		fmt.Println(err)
		return
	}
	<-doneC
}
