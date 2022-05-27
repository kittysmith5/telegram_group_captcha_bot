package msg_handler

import (
	api "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
)

func getSurfaceName(user api.User) (fullName string) {
	if user.FirstName != "" && user.LastName != "" {
		return user.FirstName + " " + user.LastName
	} else if user.LastName == "" {
		return user.FirstName
	} else {
		return strconv.Itoa(int(user.ID))
	}
}

func getUserName(user api.User) string {
	if user.UserName != "" {
		return user.UserName
	} else {
		return getSurfaceName(user)
	}
}
