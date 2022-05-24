package msg_handler

import api "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func canManageGrp(botSelf api.ChatMember) bool {
    status := botSelf.Status
    if status == "administrator" {
        canRestrictMembers := botSelf.CanRestrictMembers
        canDeleteMessages := botSelf.CanDeleteMessages
        canManageChat := botSelf.CanManageChat
        if !canRestrictMembers || !canDeleteMessages || !canManageChat {
            return false
        }
    }
    return true
}

func canSendMsg(botSelf api.ChatMember) bool {
    status := botSelf.Status
    if status == "restricted" {
        canSendMessages := botSelf.CanSendMessages
        if !canSendMessages {
            return false
        }
    }
    return true
}

func isAdmin(chatID, userID int64) bool {
    chatConf := api.ChatConfig{
        ChatID:             chatID,
        SuperGroupUsername: "",
    }
    adminConf := api.ChatAdministratorsConfig{ChatConfig: chatConf}
    admins, err := bot.GetChatAdministrators(adminConf)
    if err != nil {
        return false
    }
    for _, admin := range admins {
        if admin.User.ID == userID && canManageGrp(admin) {
            //println("=================bot is admin ================")
            return true
        }
    }
    return false
}
