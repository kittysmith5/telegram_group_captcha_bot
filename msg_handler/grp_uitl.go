package msg_handler

import (
	api "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"time"
	"v5tgbot/util"
)

func banMember(chatID, userID, sec int64) {
	if sec < 0 {
		sec = 9999999999999
	}
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
		UntilDate:        time.Now().Unix() + sec,
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

func newMembersIntoGrp(update *api.Update) {
	if update.Message == nil || !isAdmin(update.Message.Chat.ID, bot.Self.ID) {
		return
	}
	newChatMembers := update.Message.NewChatMembers
	cid := update.Message.Chat.ID
	fid := update.Message.From.ID
	member := *new(api.User)
	for _, newMember := range newChatMembers {
		member = newMember
		break
	}
	if member.IsBot && !isCreator(cid, fid) || !member.IsBot && !isAdmin(cid, fid) {
		banMember(cid, member.ID, -1)
		res, sentMsg := sendCaptcha(update, member)
		verifyMap[util.NewUUIDStr()] = VerifyType{
			newUser: member,
			res:     res,
			gid:     cid,
			mid:     sentMsg.MessageID,
		}
		time.Sleep(time.Second * 95)
		for id, verifyUser := range verifyMap {
			delMsg(cid, verifyUser.mid)
			kickMember(cid, verifyUser.newUser.ID, -1)
			delete(verifyMap, id)
		}
	}
}

func isGrp(update *api.Update) bool {
	if update.MyChatMember != nil {
		return isGrpType(update.MyChatMember.Chat.Type)
	} else if update.Message != nil {
		return isGrpType(update.Message.Chat.Type)
	} else if update.CallbackQuery != nil {
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

func isCreator(chatID, userID int64) bool {
	//println("======================")
	//println(userID)
	chatConf := api.ChatConfig{
		ChatID:             chatID,
		SuperGroupUsername: "",
	}
	adminConf := api.ChatAdministratorsConfig{ChatConfig: chatConf}
	admins, err := bot.GetChatAdministrators(adminConf)
	if err != nil {
		log.Println("Bot can't get admin of chat! There is the error: " + err.Error())
		//return true
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
