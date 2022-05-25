package msg_handler

import (
	api "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"math/rand"
	"strconv"
	"time"
)

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

func sendCapture(update *api.Update, newMember api.User) (res string, sentMsg api.Message) {
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
	msg := api.NewMessage(chatID, "123")

	msg.Text = "@" + getUserName(newMember) + "\n\n"
	msg.Text += "请在120秒内完成验证，否则永久不能入群！\n\n"
	msg.Text += "计算下面一道数学题\n\n" + txt + "\n\n\n"
	//newUser = newMember

	msg.ReplyMarkup = markup
	sentMsg, err := bot.Send(msg)
	if err != nil {
		log.Println(err)
	}
	return
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
