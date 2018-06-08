package main

import (
        "fmt"

        "github.com/yeweishuai/ctrade/okex"
)

func Printn(format string, s ...interface{}) {
    fmt.Print(format + "\n", s)
}

func main() {
    okex, err := okex.NewMarket()

    if err != nil {
        Printn("error:%s", err.Error())
        return
    }
    Printn("get okex[%v]", okex)
}
