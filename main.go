package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var token string = "7083989788:AAF-y2yluWnv0b-J5nY1fRcSoaEmXTlxk1A"
var originID int64 = -1002403516543

//var originChatId string =

func main() {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	//Начинаем обрабатывать входящее обновление. Структура update содержит информацию о полученном обновлении
	for update := range updates {
		if update.Message == nil { //Игнорируем сообщение, если текст сообщения пустой
			continue
		}

		if update.Message.ForwardFrom != nil || update.Message.ForwardFromChat != nil {
			log.Println("Получено сообщение от: ", update.Message.ForwardFrom)
		} else {
			//Выводим полезную информацию об обновлении на экран
			//NewMessageLog(update)
		}
		//Создаём заготовку сообщения, чтобы правильно её обработать
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		if update.Message.IsCommand() { //Если сообщение является командой
			switch update.Message.Command() {
			case "help":
				msg.Text = "I understand /sayhi and /status."
			case "sayhi":
				msg.Text = "Hi!"
			case "status":
				msg.Text = "I am working right now!"
			default:
				msg.Text = "I am sorry, but I don`t understand you"
			}
		}

		if update.Message != nil && !update.Message.IsCommand() { //Если сообщение является командой
			msg.Text = "I can`t understand you right now!"
		}

		//Отправляем сформированной сообщение
		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
		//NewReplyLog(msg, update)
	}
}
