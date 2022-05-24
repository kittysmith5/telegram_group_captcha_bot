package msg_handler

import "log"

import api "github.com/go-telegram-bot-api/telegram-bot-api/v5"

var bot *api.BotAPI

func botApi(botToken string) *api.BotAPI {
    botEntity, err := api.NewBotAPI(botToken)
    if err != nil {
        log.Panic("It can't get the connection with bot")
    }
    botEntity.Debug = true
    //bot.Self == getMe()
    log.Printf("Authorized on account %s\t Id: %d", botEntity.Self.UserName, botEntity.Self.ID)

    return botEntity
}

func Boot(botToken string) {
    bot = botApi(botToken)
    //消息的偏移量？
    u := api.NewUpdate(0)
    //获取更新的延迟
    u.Timeout = 6000

    updates := bot.GetUpdatesChan(u)

    for update := range updates {

        //if update.Message == nil && update.CallbackQuery == nil && update.ChatMember != nil {
        //    continue
        //}
        if isGrp(&update) {
            if update.MyChatMember != nil {
                go myChatMemberHandler(&update)
            }
            if update.Message != nil && isAdmin(update.Message.Chat.ID, bot.Self.ID) {
                go updateMsgHandler(&update)
            }
            if update.CallbackQuery != nil && isAdmin(update.CallbackQuery.Message.Chat.ID, bot.Self.ID) {
                go callbackQueryHandler(&update)
            }
        } else {
            continue
        }
    }
}
