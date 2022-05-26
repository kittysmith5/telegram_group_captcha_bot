package main

import (
	"flag"
	"fmt"
	"v5tgbot/msg_handler"
)

func main() {
	//token := flag.String("t", "", "token of your bot")
	token := flag.String("t", "", "token of your bot")
	flag.Parse()
	fmt.Println("YOUR TOKEN IS: " + *token)
	msg_handler.Boot(*token)
}
