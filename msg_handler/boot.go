package msg_handler

import (
	api "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

var bot *api.BotAPI

func botApi(botToken string) *api.BotAPI {
	botEntity, err := api.NewBotAPI(botToken)
	if err != nil {
		log.Panic("It can't get the connection with bot")
	}
	//botEntity.Debug = true
	//bot.Self == getMe()
	log.Printf("Authorized on account %s\t Id: %d", botEntity.Self.UserName, botEntity.Self.ID)
	return botEntity
}

func Boot(botToken string) {
	bot = botApi(botToken)
	//消息的偏移量？-1表示只保留最后一条update
	u := api.NewUpdate(-1)
	//Timeout in seconds for long polling
	//u.Timeout = 0
	updates := bot.GetUpdatesChan(u)
	//get the last update
	//begin := true
	for update := range updates {
		//if begin {
		//    begin = false
		//    continue
		//}
		print(update.UpdateID)
		go updateHandler(&update)
	}
}

func updateHandler(update *api.Update) {
	if update.MyChatMember != nil {
		myChatMemberHandler(update)
	} else if update.Message != nil {
		updateMsgHandler(update)
	} else if update.CallbackQuery != nil {
		callbackQueryHandler(update)
	}
}
