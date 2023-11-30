package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/hamidteimouri/gommon/htenvier"
	"log"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(htenvier.Env("BOT_TOKEN"))
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore non-Message updates
			continue
		}

		/*msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

		switch update.Message.Text {
		case "open":
			msg.ReplyMarkup = numericKeyboard
		case "close":
			msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)

		}*/

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		if update.Message.IsCommand() {
			// Extract the command from the Message.
			switch update.Message.Command() {
			case "start":
				msg.Text = getMsgWelcome(update.Message.Chat.FirstName)
				msg.ReplyMarkup = mainMenu
			case "back":
				msg.Text = "I understand /sayhi and /status."
				msg.ReplyMarkup = mainMenu
			case "sayhi":
				msg.Text = "Hi :)"
			default:
				msg.Text = "در حال پیاده سازی"
			}
		} else {
			// Extract the command from the Message.
			switch update.Message.Text {
			case BtnHome:
				msg.Text = getMsgBackToHome()
				msg.ReplyMarkup = mainMenu
			case BtnTetherPrice:
				msg.Text = getMsgTetherPrice()
				msg.ReplyMarkup = mainMenuWithBack
			case BtnCalculator:
				msg.Text = getMsgSelectYourCoin()
				msg.ReplyMarkup = coinsMenuWithBack
			default:
				msg.Text = "در حال پیاده سازی امکانات"
				msg.ReplyMarkup = mainMenu
			}
		}

		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}
