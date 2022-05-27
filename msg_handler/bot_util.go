package msg_handler

import (
	api "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func botCanManageGrp(botSelf api.ChatMember) bool {
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

func botCanSendMsg(botSelf api.ChatMember) bool {
	status := botSelf.Status
	if status == "restricted" {
		canSendMessages := botSelf.CanSendMessages
		if !canSendMessages {
			return false
		}
	}
	return true
}
