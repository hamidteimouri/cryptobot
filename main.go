package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/hamidteimouri/gommon/htenvier"
	"log"
	"strconv"
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
			case BtnBalanceUsdtTrc20:
				msg.Text = getMsgEnterYourWalletAddress()
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
func Format(n int64) string {
	in := strconv.FormatInt(n, 10)
	numOfDigits := len(in)
	if n < 0 {
		numOfDigits-- // First character is the - sign (not a digit)
	}
	numOfCommas := (numOfDigits - 1) / 3

	out := make([]byte, len(in)+numOfCommas)
	if n < 0 {
		in, out[0] = in[1:], '-'
	}

	for i, j, k := len(in)-1, len(out)-1, 0; ; i, j = i-1, j-1 {
		out[j] = in[i]
		if i == 0 {
			return string(out)
		}
		if k++; k == 3 {
			j, k = j-1, 0
			out[j] = ','
		}
	}
}
