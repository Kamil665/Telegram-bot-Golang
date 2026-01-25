package services

import (
	"fmt"

	"telegram-NewBot/keyboards"
	"telegram-NewBot/repositories"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func StartCommand(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	text := "Hello! I'm your task manager bot."
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
	msg.ReplyMarkup = keyboards.CmdKeyboard()
	if _, err := bot.Send(msg); err != nil {
		panic(err)
	}
}

var userCoins = make(map[int64]int)

func MoneyTask(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	userID := update.Message.From.ID

	userCoins[userID] += 1

	text := fmt.Sprintf("+1 money 💸. You have %d coins!", userCoins[userID])

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
	if _, err := bot.Send(msg); err != nil {
		panic(err)
	}
}

func SetTaskCallback(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	if update.CallbackQuery != nil {
		callback := update.CallbackQuery

		data := callback.Data
		chatID := callback.Message.Chat.ID
		messageID := callback.Message.MessageID

		msg := tgbotapi.NewEditMessageText(chatID, messageID, "You selected: "+data)
		_, err := bot.Send(msg)
		if err != nil {
			panic(err)
		}

		answer := tgbotapi.NewCallback(callback.ID, "Callback received!")
		_, err = bot.Request(answer)
		if err != nil {
			panic(err)
		}
	}
}

func SaveTaskToDatabase(bot *tgbotapi.BotAPI, update tgbotapi.Update, taskText string) {
	task, err := repositories.CreateTask(update.Message.From.ID, taskText)
	if err == nil && task != nil {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Task saved!"+taskText)
		bot.Send(msg)
	}
}
