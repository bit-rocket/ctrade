package main

import (
        "fmt"

        "github.com/yeweishuai/ctrade/okex"
)

func main() {
    okex, err := okex.NewMarket()
    printn := fmt.Println

    if err != nil {
        Printn("error:%s", err.Error())
        return
    }
    printn("get okex[%v]\n", okex)
}
