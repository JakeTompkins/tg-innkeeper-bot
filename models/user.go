package models

const USER_TABLE = "users"

type User struct {
	Id              int    `sql:"id"`
	TelegramId      int    `sql:"telegram_id"`
	DefaultTimezone string `sql:"default_timezone"`
}

func NewUser(id int, telegramId int, defaultTimezone string) *User {
	return &User{Id: id, TelegramId: telegramId, DefaultTimezone: defaultTimezone}
}
