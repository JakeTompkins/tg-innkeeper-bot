package main

import (
	"tg-group-scheduler/telegram"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	bot := telegram.NewTelegramBot()
	bot.Listen()
}
