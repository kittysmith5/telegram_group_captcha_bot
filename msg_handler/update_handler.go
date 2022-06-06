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

var counterMap = map[int64]int8{}

func updateMsgHandler(update *api.Update) {
	if isGrp(update) && update.Message != nil {
		upMsg := update.Message
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
			userIsRight := verifyType.newUser.ID == update.CallbackQuery.From.ID
			resIsRight := verifyType.res == update.CallbackQuery.Data
			msgIsRight := verifyType.mid == mid
			cidISRight := verifyType.gid == chatID

			if userIsRight && resIsRight && msgIsRight && cidISRight {
				delMsg(chatID, mid)
				unRestrictMember(chatID, verifyType.newUser.ID)
				delete(verifyMap, id)
				sendAnswerCallBack(update.CallbackQuery.ID, "    恭喜，你通过了验证!")
			} else if userIsRight && msgIsRight && cidISRight {
				sendAnswerCallBack(update.CallbackQuery.ID, "对不起，回答错误，请在6个小时后重新加群！")
				delMsg(chatID, verifyType.mid)
				tipMsg := getUserName(verifyType.newUser) +
					"\n对不起，回答错误，请在6个小时后重新加群！如果机器人误操作，请联系群管理员！"
				answerTipMsg := sendMarkDownMsg(chatID, tipMsg)
				time.Sleep(time.Second * 30)
				kickMember(chatID, verifyType.newUser.ID, 3600*6)
				delete(verifyMap, id)
				delMsg(chatID, answerTipMsg.MessageID)
			} else if !userIsRight && cidISRight && !isAdmin(chatID, update.CallbackQuery.From.ID) {
				wrongUserID := update.CallbackQuery.From.ID
				counterMap[wrongUserID] += 1
				sendAnswerCallBack(update.CallbackQuery.ID, "不要乱点他人验证消息, 多次点击将会禁言!")
				for userID, counter := range counterMap {
					if counter == 3 {
						banMember(chatID, userID, 900)
						atName := getUserName(*update.CallbackQuery.From)
						txt := atName + "\n" + "\n乱点他人验证消息三次，禁言15分钟！"
						sentMsg := sendMarkDownMsg(chatID, txt)
						counterMap[userID] = 0
						time.Sleep(time.Second * 10)
						delMsg(chatID, sentMsg.MessageID)
					}
				}
			}
		}
	}
}

func myChatMemberHandler(update *api.Update) {
	newUser := update.MyChatMember.NewChatMember
	newUserID := newUser.User.ID
	if newUserID == bot.Self.ID {
		if !botCanManageGrp(newUser) {
			cid := update.MyChatMember.Chat.ID
			sentTxtMsg := sendTxtMsg(cid,
				"请给我删除消息权限，管理群权限，禁言删除群成员权限（管理员权限中设置）！才能进行入群验证")
			time.Sleep(time.Second * 30)
			delMsg(cid, sentTxtMsg.MessageID)
		}
		if !botCanSendMsg(newUser) && !update.MyChatMember.From.IsBot {
			//priChatID :=
			userID := update.MyChatMember.From.ID
			atUser := getUserName(update.MyChatMember.From)
			grpTitle := update.MyChatMember.Chat.Title
			grpName := update.MyChatMember.Chat.UserName
			grpLink := "[" + grpTitle + "](t.me/" + grpName + ")"

			sentMDMsg := sendMarkDownMsg(userID,
				atUser+"\n请给我\t"+grpLink+"\t群\n发送消息权限（机器人权限设置）")
			time.Sleep(time.Second * 300)
			delMsg(userID, sentMDMsg.MessageID)
		}
	}
}
