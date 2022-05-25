package main

import (
	"flag"
	"fmt"
	"v5tgbot/msg_handler"
)

func main() {
	token := flag.String("t", "", "token of your bot")
	flag.Parse()
	fmt.Println(*token)
	msg_handler.Boot(*token)
}
