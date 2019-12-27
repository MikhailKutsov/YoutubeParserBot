package main

import (
	"encoding/json"
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"golang.org/x/net/context"
	"io/ioutil"
	"log"
	"os"
)

const youtubeConfig = "client_secret.json"
const telegramConfig = "config.json"
//const playlistId = "PL1eQAFFC9DaMS2zQBiLxW8EPcoYCkzZug"

var playlistId = ""

var youtubeService = GetYoutubeService(context.Background(), getYoutubeConfig(youtubeConfig))

type Config struct {
	TelegramBotToken string
}

func getYoutubeConfig(path string) []byte  {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}
	return b
}

func getTGConfig(path string) Config {
	file, _ := os.Open(path)
	decoder := json.NewDecoder(file)
	configuration := Config{}
	err := decoder.Decode(&configuration)
	if err != nil {
		log.Panic(err)
	}

	return configuration
}



func main() {
	bot := getTGBot(getTGConfig(telegramConfig))

	// инициализируем канал, куда будут прилетать обновления от API
	var ucfg tgbotapi.UpdateConfig = tgbotapi.NewUpdate(0)
	ucfg.Timeout = 60
	//updates, err := bot.GetUpdates(ucfg)
	updates, err := bot.GetUpdatesChan(ucfg)
	if err != nil {
		log.Panic(err)
	}
	// читаем обновления из канала
	for {
		select {
		case update := <-updates:
			// Пользователь, который написал боту
			UserName := update.Message.From.UserName
			log.Printf("[%s] %d %s", update.Message.Text)

			// ID чата/диалога.
			// Может быть идентификатором как чата с пользователем
			// (тогда он равен UserID) так и публичного чата/канала
			ChatID := update.Message.Chat.ID

			// Текст сообщения
			Text := update.Message.Text

			log.Printf("[%s] %d %s", UserName, ChatID, Text)

			// Ответим пользователю его же сообщением
			//reply := parse(update.Message.Text)
		/*	switch Text {

			case "/playlistId"

			}*/
			parsed := parse(update.Message.Text)
			msg := tgbotapi.NewMessage(ChatID, parsed)
			bot.Send(msg)
			if parsed != "" {
				reply := playlistItemInsert(youtubeService, parsed, playlistId)
				msg := tgbotapi.NewMessage(ChatID, reply)
				bot.Send(msg)
			}
		}
	}
}