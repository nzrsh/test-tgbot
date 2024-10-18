package main

import (
	"fmt"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func NewMessageLog(u tgbotapi.Update) {
	if u.Message != nil { // Проверяем, что это сообщение, а не другой тип обновления
		// Получаем имя пользователя (никнейм)
		senderName := u.Message.From.FirstName
		if u.Message.From.LastName != "" {
			senderName += " " + u.Message.From.LastName
		}

		// Логин отправителя, если указан
		senderUsername := u.Message.From.UserName
		if senderUsername == "" {
			senderUsername = "Не указан"
		}

		// Текст сообщения или команда
		var messageText string
		if u.Message.IsCommand() {
			messageText = "Команда: " + u.Message.Command()
		} else {
			messageText = "Сообщение: " + u.Message.Text
		}

		// Chat ID
		chatID := u.Message.Chat.ID

		// Время отправки
		messageTime := time.Unix(int64(u.Message.Date), 0).Format(time.RFC1123)

		// Выводим информацию на экран

		fmt.Printf(
			"======NEW MESSAGE======\nОтправитель: %s\n%s\nЧат: %d\nВремя: %s\n",
			senderUsername,
			messageText,
			chatID,
			messageTime,
		)
	} else {
		fmt.Println("Обновление не содержит сообщения.")
	}
}

func NewReplyLog(msg tgbotapi.MessageConfig, u tgbotapi.Update) {
	fmt.Printf(
		"=======NEW REPLY=======\nПолучатель: %s\nЧат: %d\nТекст ответа: %s\nВремя: %s\n",
		u.Message.From.UserName,
		msg.ChatID,
		msg.Text,
		time.Now().Format(time.RFC1123),
	)
}
