package main

import (
    "flag"
    "fmt"
    "v5tgbot/msg_handler"
)

var testMap = map[int]int{}

func main() {
    token := flag.String("t", "1800998568:AAEWA8YZk_qVugx2ryBq2aoVUV6H7Vn6_gE", "token of your bot")
    flag.Parse()
    fmt.Println(*token)
    msg_handler.Boot(*token)
}
