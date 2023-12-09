package main

import (
	"fmt"
	bt "github.com/SakoDroid/telego/v2"
	cfg "github.com/SakoDroid/telego/v2/configs"
	objs "github.com/SakoDroid/telego/v2/objects"
	"github.com/hamidteimouri/gommon/htenvier"
	"strconv"
)

func main() {
	apiToken := htenvier.Env("BOT_TOKEN")
	bot, err := bt.NewBot(cfg.Default(apiToken))

	if err != nil {
		panic(err)
	}

	// The general update channel.
	updateChannel := *(bot.GetUpdateChannel())

	//Adding a handler. Everytime the bot receives message "hi" in a private chat, it will respond "hi to you too".

	bot.AddHandler("/start", func(u *objs.Update) {

		//Create the custom keyboard.
		kb := bot.CreateKeyboard(false, false, false, false, "nemidonam")

		//Add buttons to it. First argument is the button's text and the second one is the row number that the button will be added to it.
		kb.AddButton("button1", 1)
		kb.AddButton("button2", 1)
		kb.AddButton("button3", 2)
		kb.AddButton("button3", 2)
		kb.AddButton("button3", 2)

		//Sends the message along with the keyboard.
		_, err := bot.AdvancedMode().ASendMessage(u.Message.Chat.Id, "hi to you too", "", u.Message.MessageId, 0, false, false, nil, false, false, kb)
		if err != nil {
			fmt.Println(err)
		}
	}, "private")

	// Monitores any other update. (Updates that don't contain text message "hi" in a private chat)
	go func() {
		for {
			update := <-updateChannel
			fmt.Println(update.Update_id)

			//Some processing on the update
		}
	}()

	bot.Run(true)

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
