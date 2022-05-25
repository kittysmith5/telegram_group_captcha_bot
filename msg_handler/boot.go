package msg_handler

import (
	"log"
)

import api "github.com/go-telegram-bot-api/telegram-bot-api/v5"

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
	//消息的偏移量？
	u := api.NewUpdate(-1)
	//获取更新的延迟
	u.Timeout = 0
	updates := bot.GetUpdatesChan(u)

	//get the last update
	begin := true
	for update := range updates {
		if begin {
			begin = false
			continue
		}
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
