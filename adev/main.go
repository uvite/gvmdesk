package main

import (
	"fmt"
	"github.com/uvite/gvmdesk/gvmbot"

	"os"
)

func main() {
	pwd, _ := os.Getwd()

	filepath := fmt.Sprintf("%s/%s", pwd, ".env.local")
	configpath := fmt.Sprintf("%s/%s", pwd, "bbgo.yaml")

	ex := gvmbot.New(filepath, configpath, "abc")
	//a, e := ex.GetAccount()
	//fmt.Println(a, e)
	//
	e := ex.SubKline()
	fmt.Println(e)

}
