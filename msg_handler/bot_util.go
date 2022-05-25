package msg_handler

import (
	api "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func canManageGrp(botSelf api.ChatMember) bool {
	status := botSelf.Status
	if status == "member" {
		return false
	} else if status == "administrator" {
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
func isCreator(chatID, userID int64) bool {
	chatConf := api.ChatConfig{
		ChatID:             chatID,
		SuperGroupUsername: "",
	}
	adminConf := api.ChatAdministratorsConfig{ChatConfig: chatConf}
	admins, err := bot.GetChatAdministrators(adminConf)
	if err != nil {
		log.Println("Bot can't get admin of chat! There is the error: " + err.Error())
		return false
	}
	for _, admin := range admins {
		if admin.Status == "creator" && admin.User.ID == userID {
			return true
		}
	}
	return false
}

func isAdmin(chatID, userID int64) bool {
	chatConf := api.ChatConfig{
		ChatID:             chatID,
		SuperGroupUsername: "",
	}
	adminConf := api.ChatAdministratorsConfig{ChatConfig: chatConf}
	admins, err := bot.GetChatAdministrators(adminConf)
	if err != nil {
		log.Println("Bot can't get admin of chat! There is the error: " + err.Error())
		return false
	}
	for _, admin := range admins {
		if admin.User.ID == userID /*&& canManageGrp(admin)*/ {
			return true
		}
	}
	return false
}
