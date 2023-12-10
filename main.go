package main

import (
	"context"
	"fmt"
	bt "github.com/SakoDroid/telego/v2"
	cfg "github.com/SakoDroid/telego/v2/configs"
	objs "github.com/SakoDroid/telego/v2/objects"
	"github.com/hamidteimouri/gommon/htenvier"
	"github.com/hamidteimouri/tronscansdk"
	"github.com/sirupsen/logrus"
	"regexp"
	"strconv"
)

const (
	ChatTypePrivate = "private"
)

func main() {
	//var err error

	apiToken := htenvier.Env("BOT_TOKEN")
	bot, err := bt.NewBot(cfg.Default(apiToken))

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Panic("failed to initiate bot")
	}

	// The general update channel.
	updateChannel := *(bot.GetUpdateChannel())

	_ = bot.AddHandler("/start", func(u *objs.Update) {
		kb := bot.CreateKeyboard(true, false, false, true, "nemidonam")
		kb.AddButton(BtnTetherPrice, 1)
		kb.AddButton(BtnBalanceUsdtTrc20, 1)

		_, err = bot.AdvancedMode().ASendMessage(u.Message.Chat.Id, getMsgWelcome(u.Message.Chat.FirstName), "",
			u.Message.MessageId, 0,
			false, false, nil, false, false, kb)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"err": err,
			}).Error("failed to send /start message")
		}
	}, ChatTypePrivate)

	_ = bot.AddHandler(BtnTetherPrice, func(u *objs.Update) {

		// Sends the message along with the keyboard.
		_, err = bot.SendMessage(u.Message.Chat.Id, getMsgTetherPrice(), "",
			0, false,
			false)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"err": err,
			}).Error("failed to send btn-tether-price message")
		}
	}, ChatTypePrivate)

	_ = bot.AddHandler(BtnBalanceUsdtTrc20, func(u *objs.Update) {

		// Sends the message along with the keyboard.
		_, err = bot.SendMessage(u.Message.Chat.Id, getMsgEnterYourWalletAddress(), "",
			0, false,
			false)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"err": err,
			}).Error("failed to send btn-balance-wallet-trc20 message")
		}
	}, ChatTypePrivate)

	go func() {
		for {
			update := <-updateChannel
			fmt.Println(update.Update_id)
			fmt.Println(update.Message.Text)

			if IsTronAddress(update.Message.Text) {
				usdt, err := tronscansdk.GetBalanceOfUsdt(context.Background(), update.Message.Text)
				msg := ""
				if err != nil {
					msg = "خطایی رخ داده است"
					logrus.WithFields(logrus.Fields{
						"err": err,
					}).Error("failed to get balance of wallet")
				} else {
					msg = usdt
				}

				bot.SendMessage(update.Message.Chat.Id, msg, "", update.Message.MessageId, false, false)
			}

			//Some processing on the update
		}
	}()

	err = bot.Run(true)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Error("failed to run bot")
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

const (
	TronTrc20Pattern = "^T[1-9A-HJ-NP-Za-km-z]{33}$"
)

var (
	TronRegex = regexp.MustCompile(TronTrc20Pattern)
)

func IsTronAddress(text string) bool {
	return TronRegex.MatchString(text)
}
