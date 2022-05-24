package msg_handler

import api "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func getUserName(user api.User) (username string) {
    if user.UserName != "" {
        username = user.UserName
        return
    } else {
        username = user.FirstName + " " + user.LastName
        return
    }
}
