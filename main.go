package main

import (
	"fmt"
	"tg-group-scheduler/telegram"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	bot := telegram.NewTelegramBot()
	err := bot.Listen()

	if err != nil {
		fmt.Println("Bot listening successfully")
	} else {
		fmt.Println("Bot failed to listen")
	}
}
