package msg_handler

import (
	api "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func wrapUserLink(user api.User) (msgTxt string) {
	fullName := getSurfaceName(user)
	msgTxt = "[" + fullName + "]" + "(tg://user?id=" + strconv.Itoa(int(user.ID)) + ")"
	return
}

func sendTxtMsg(cid int64, txt string) (sentMsg api.Message) {
	msg := api.NewMessage(cid, txt)
	sentMsg, err := bot.Send(msg)
	if err != nil {
		log.Println("发送文字消息错误： " + err.Error())
	}
	return
}

func sendMarkDownMsg(cid int64, txt string) (sentMsg api.Message) {
	msg := api.NewMessage(cid, txt)
	msg.ParseMode = "Markdown"
	sentMsg, err := bot.Send(msg)
	if err != nil {
		log.Println("发送Markdown消息错误： " + err.Error())
	}
	return
}

func sendCaptcha(update *api.Update, newMember api.User) (res string, sentMsg api.Message) {
	plusSticker := "\u2795"
	chatID := update.Message.Chat.ID
	//newMembers := *update.Message.NewChatMembers
	//设置随机种子
	rand.Seed(time.Now().UnixNano())
	randomNum1 := rand.Int()%250 + 200
	randomNum2 := rand.Int()%250 + 200
	randomRes := randomNum1 + randomNum2
	txt := int2Sticker(randomNum1) + "\t " + plusSticker + "\t " + int2Sticker(randomNum2)
	txt += "等于多少" + "\u2753"

	markA := strconv.Itoa(randomRes + 10)
	markB := strconv.Itoa(randomRes - 10)
	markC := strconv.Itoa(randomRes + 20)
	markD := strconv.Itoa(randomRes - 20)
	markE := strconv.Itoa(randomRes + 30)

	res = strconv.Itoa(randomRes)

	rows := make([]api.InlineKeyboardButton, 5)
	randomMark := rand.Int() % 5

	rows[0] = api.NewInlineKeyboardButtonData(markA, markA)
	rows[1] = api.NewInlineKeyboardButtonData(markB, markB)
	rows[2] = api.NewInlineKeyboardButtonData(markC, markC)
	rows[3] = api.NewInlineKeyboardButtonData(markD, markD)
	rows[4] = api.NewInlineKeyboardButtonData(markE, markE)

	rows[randomMark] = api.NewInlineKeyboardButtonData(res, res)
	markup := api.NewInlineKeyboardMarkup(rows)

	msg := api.NewMessage(chatID, "")
	msg.Text = getUserName(newMember) + "\n\n\n"
	msg.Text += "待验证者：" + wrapUserLink(newMember)
	//println(getUserName(newMember))
	msg.Text += "\n\n请在120秒内完成验证，否则永久不能入群！\n\n" +
		"请计算下面一道数学题\n\n\n\n" + txt + "\n\n"
	// __ is markdown char, so it needs to escape with `\_`
	msg.Text = strings.Replace(msg.Text, "_", `\_`, -1)
	msg.ParseMode = "MarkdownV2"
	msg.ReplyMarkup = markup
	//sendTxtMsg(chatID, "@"+getUserName(newMember)+"\n")
	//sendMarkDownMsg(chatID,)
	sentMsg, err := bot.Send(msg)
	if err != nil {
		log.Println(err)
	}
	return
}

func sendAnswerCallBack(callbackQueryID, txt string) {
	var callbackConfig = api.CallbackConfig{
		CallbackQueryID: callbackQueryID,
		Text:            txt,
		ShowAlert:       true,
		URL:             "",
		CacheTime:       0,
	}
	_, _ = bot.Send(callbackConfig)
	//if err != nil {
	//	return
	//}
}

func int2Sticker(num int) (str string) {
	str = ""
	stickerSuffix := "\ufe0f\u20e3"
	for {
		if num == 0 {
			break
		}
		reminder := num % 10
		str = strconv.Itoa(reminder) + stickerSuffix + str
		num /= 10
	}
	return
}
