package msg_handler

import (
	api "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"time"
	"v5tgbot/util"
)

func banNewMember(chatID, userID int64) {
	memberConfig := api.ChatMemberConfig{
		ChatID:             chatID,
		SuperGroupUsername: "",
		ChannelUsername:    "",
		UserID:             userID,
	}

	blankPermissions := api.ChatPermissions{
		CanSendMessages:       false,
		CanSendMediaMessages:  false,
		CanSendPolls:          false,
		CanSendOtherMessages:  false,
		CanAddWebPagePreviews: false,
		CanChangeInfo:         false,
		CanInviteUsers:        false,
		CanPinMessages:        false,
	}

	restrictConfig := api.RestrictChatMemberConfig{
		ChatMemberConfig: memberConfig,
		UntilDate:        time.Now().Unix() + 9999999999999,
		Permissions:      &blankPermissions,
	}
	_, _ = bot.Send(restrictConfig)
}

func kickMember(chatID, userID, sec int64) {
	if sec < 0 {
		sec = 9999999999999
	}
	memberConf := api.ChatMemberConfig{
		ChatID:             chatID,
		SuperGroupUsername: "",
		ChannelUsername:    "",
		UserID:             userID,
	}
	kickChatMemberConf := api.KickChatMemberConfig{
		ChatMemberConfig: memberConf,
		UntilDate:        time.Now().Unix() + sec,
		RevokeMessages:   false,
	}
	_, _ = bot.Send(kickChatMemberConf)
}

/*func unbanMember(chatID, userID int64) {
    memberConf := api.ChatMemberConfig{
        ChatID:             chatID,
        SuperGroupUsername: "",
        ChannelUsername:    "",
        UserID:             userID,
    }
    unbanConf := api.UnbanChatMemberConfig{
        ChatMemberConfig: memberConf,
        OnlyIfBanned:     false,
    }
    _, _ = bot.Send(unbanConf)
}*/

func unRestrictMember(chatID, userID int64) {
	memberConfig := api.ChatMemberConfig{
		ChatID:             chatID,
		SuperGroupUsername: "",
		ChannelUsername:    "",
		UserID:             userID,
	}

	allPermissions := api.ChatPermissions{
		CanSendMessages:       true,
		CanSendMediaMessages:  true,
		CanSendPolls:          true,
		CanSendOtherMessages:  true,
		CanAddWebPagePreviews: true,
		CanChangeInfo:         true,
		CanInviteUsers:        true,
		CanPinMessages:        true,
	}

	restrictConfig := api.RestrictChatMemberConfig{
		ChatMemberConfig: memberConfig,
		UntilDate:        time.Now().Unix() + 9999999999999,
		Permissions:      &allPermissions,
	}
	_, _ = bot.Send(restrictConfig)
}

func botItselfIntoGrp(update *api.Update) {
	newChatMembers := update.Message.NewChatMembers
	for _, newMember := range newChatMembers {
		if newMember.IsBot && newMember.ID == bot.Self.ID {
			cid := update.Message.Chat.ID
			sendTxtMsg(cid, "请给我管理员权限，才能开启入群验证功能！")
			return
		}
	}
}

func newMembersIntoGrp(update *api.Update) {
	newChatMembers := update.Message.NewChatMembers
	cid := update.Message.Chat.ID
	for _, newMember := range newChatMembers {
		println(newMember.ID)
		banNewMember(cid, newMember.ID)
		res, sentMsg := sendCapture(update, newMember)

		verifyMap[util.NewUUIDStr()] = VerifyType{
			newUser: newMember,
			res:     res,
			gid:     cid,
			mid:     sentMsg.MessageID,
		}
	}
}

func isGrp(update *api.Update) bool {
	if update.MyChatMember != nil {
		return isGrpType(update.MyChatMember.Chat.Type)
	}
	if update.Message != nil {
		return isGrpType(update.Message.Chat.Type)
	}
	if update.CallbackQuery != nil {
		return isGrpType(update.CallbackQuery.Message.Chat.Type)
	}
	return false
}

func isGrpType(typeStr string) bool {
	if typeStr == "supergroup" || typeStr == "group" {
		return true
	} else {
		return false
	}
}

func delMsg(chatID int64, msgID int) {
	delMsgConf := api.DeleteMessageConfig{
		ChannelUsername: "",
		ChatID:          chatID,
		MessageID:       msgID,
	}
	_, _ = bot.Send(delMsgConf)
}
