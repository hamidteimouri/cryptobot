package main

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

var mainMenu = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(BtnLivePriceWithText),
		tgbotapi.NewKeyboardButton(BtnLivePriceWithPic),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(BtnTetherPrice),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(BtnCalculator),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(BtnEvents),
		tgbotapi.NewKeyboardButton(BtnFear),
	),
)

var mainMenuWithBack = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(BtnHome),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(BtnLivePriceWithText),
		tgbotapi.NewKeyboardButton(BtnLivePriceWithPic),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(BtnTetherPrice),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(BtnCalculator),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(BtnEvents),
		tgbotapi.NewKeyboardButton(BtnFear),
	),
)
var coinsMenu = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("USDT"),
		tgbotapi.NewKeyboardButton("BTC"),
		tgbotapi.NewKeyboardButton("ETH"),
		tgbotapi.NewKeyboardButton("TRX"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("BNB"),
		tgbotapi.NewKeyboardButton("PEPE"),
		tgbotapi.NewKeyboardButton("DOGE"),
		tgbotapi.NewKeyboardButton("XRP"),
	),
)
var coinsMenuWithBack = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(BtnHome),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("USDT"),
		tgbotapi.NewKeyboardButton("BTC"),
		tgbotapi.NewKeyboardButton("ETH"),
		tgbotapi.NewKeyboardButton("TRX"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("BNB"),
		tgbotapi.NewKeyboardButton("PEPE"),
		tgbotapi.NewKeyboardButton("DOGE"),
		tgbotapi.NewKeyboardButton("XRP"),
	),
)
