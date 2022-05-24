package main

import (
    "flag"
    "fmt"
    "v5tgbot/msg_handler"
)

var testMap = map[int]int{}

func main() {
    //token := flag.String("t", "", "token of your bot")
    token := flag.String("t", "YOU_TELEGRAM_BOT_TOKEN", "token of your bot")
    flag.Parse()
    fmt.Println(*token)
    msg_handler.Boot(*token)
}
