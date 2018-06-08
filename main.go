package main

import (
        "fmt"
        "time"

        "github.com/yeweishuai/ctrade/okex"
)

func main() {
    okex, err := okex.NewMarket()
    printn := fmt.Println

    if err != nil {
        printn("error:%s", err.Error())
        return
    }
    printn("get okex[%v]\n", okex)
    time.Sleep(time.Hour)
}
