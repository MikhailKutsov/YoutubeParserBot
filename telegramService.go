package main

import (
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"log"
)

func getTGBot(config Config) *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI(config.TelegramBotToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)
	return bot
}