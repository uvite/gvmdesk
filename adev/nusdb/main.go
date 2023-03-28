package main

import (
	"context"
	gtools "github.com/uvite/gvmdesk/controller"
	"github.com/uvite/gvmdesk/internal"
)

func main() {
	a := gtools.NewApp()
	ctx := context.Background()

	a.OnStartup(ctx)
	//a.DBFile = fmt.Sprintf(configs.DBFile, "confDir")
	//a.Db = internal.NewXormEngine(a.DBFile)
	item := internal.AlertItem{
		Title:    "rsi",
		Symbol:   "LTCUSDT",
		Interval: "15m",
		Path:     "/hein/gvmdesk/js/4.js",
	}
	//a.AddAlertItem(item)

	//ab := a.GetAlertList()
	//fmt.Println(ab)
	item.Id = 4
	item.Status = true
	item.Interval = "1m"
	item.Content = `export default function () {

	 console.log(close.last())

	};`
	a.UpdateAlertItem(item)
	////fmt.Printf("%+v", ab1)
	//a.AddSymbolInterval(item.Symbol, item.Interval)
	//a.RunTestFile(item)
	//
	//cmdutil.WaitForSignal(ctx, syscall.SIGINT, syscall.SIGTERM)

}
