package msg_handler

import (
	api "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"time"
)

type VerifyType struct {
	newUser api.User
	res     string
	gid     int64
	mid     int
}

var verifyMap = map[string]VerifyType{}

func updateMsgHandler(update *api.Update) {
	if isGrp(update) && update.Message != nil {
		upMsg := update.Message
		println("\n=================msg ==================")
		if upMsg.NewChatMembers != nil {
			chatID := update.Message.Chat.ID
			delMsg(chatID, update.Message.MessageID)
			newMembersIntoGrp(update)
		} else if upMsg.LeftChatMember != nil {
			delMsg(update.Message.Chat.ID, update.Message.MessageID)
		}
	}
}

func callbackQueryHandler(update *api.Update) {
	chatID := update.CallbackQuery.Message.Chat.ID
	if isGrp(update) && isAdmin(chatID, bot.Self.ID) {
		mid := update.CallbackQuery.Message.MessageID
		for id, verifyType := range verifyMap {
			condition1 := verifyType.newUser.ID == update.CallbackQuery.From.ID
			condition2 := verifyType.res == update.CallbackQuery.Data
			condition3 := verifyType.mid == mid
			condition4 := verifyType.gid == chatID

			if condition1 && condition2 && condition3 && condition4 {
				delMsg(chatID, mid)
				unRestrictMember(chatID, verifyType.newUser.ID)
				delete(verifyMap, id)
			} else if condition1 && condition3 {
				txt := "@" + verifyType.newUser.UserName + "\n" + "对不起，回答错误，请在6个小时后重新加群！"
				delMsg(chatID, mid)
				sentMsg := sendTxtMsg(chatID, txt)
				time.Sleep(time.Second * 8)
				delMsg(chatID, sentMsg.MessageID)
				kickMember(chatID, verifyType.newUser.ID, 3600*6)
				delete(verifyMap, id)
				time.Sleep(time.Second * 5)
			}
		}
	}
}

func myChatMemberHandler(update *api.Update) {
	newUser := update.MyChatMember.NewChatMember
	newUserID := newUser.User.ID
	if newUserID == bot.Self.ID {
		if !canManageGrp(newUser) {
			cid := update.MyChatMember.Chat.ID
			sentTxtMsg := sendTxtMsg(cid, "请给我删除消息权限，管理群权限，禁言删除群成员权限（管理员权限中设置）！才能进行入群验证")
			time.Sleep(time.Second * 30)
			delMsg(cid, sentTxtMsg.MessageID)
		}
		if !canSendMsg(newUser) && !update.MyChatMember.From.IsBot {
			//priChatID :=
			userID := update.MyChatMember.From.ID
			atUser := "@" + getUserName(update.MyChatMember.From)
			grpTitle := update.MyChatMember.Chat.Title
			grpName := update.MyChatMember.Chat.UserName
			grpLink := "[" + grpTitle + "](t.me/" + grpName + ")"

			sentMDMsg := sendMarkDownMsg(userID, atUser+"\n请给我\t"+grpLink+"\t群\n发送消息权限（机器人权限设置）")
			time.Sleep(time.Second * 300)
			delMsg(userID, sentMDMsg.MessageID)
		}
	}
}
