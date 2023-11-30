package main

import (
	"context"
	"fmt"
	"time"
)

func getMsgWelcome(name string) string {
	return fmt.Sprintf("سلام %s عزیز،\nچطوری میتونم کمکت کنم؟", name)
}
func getMsgBackToHome() string {
	return "منتظرم انتخاب کنی!"
}

func getMsgSelectYourCoin() string {
	return "ارز مورد نظر خودتو انتخاب کن"
}
func getMsgEnterYourWalletAddress() string {
	return "آدرس ولت خود را وارد کنید"
}
func getMsgEnterYourAmount() string {
	return "مقدار مورد نظر خودتو وارد کن"
}
func getMsgEnterYourDestinationCoin() string {
	return "این موارد رو، به چه ارزی واست حساب کنم؟"
}

func getMsgTetherPrice() string {
	r, err := GetTetherPrice(context.Background())
	if err != nil {
		return "خطایی رخ داده است"
	}
	s1 := Format(r.Sell.Round(0).IntPart())
	s2 := Format(r.Buy.Round(0).IntPart())
	loc, _ := time.LoadLocation(TimeZoneTehran)
	n := time.Now().In(loc).Format(time.RFC3339)

	msg := `
▫️خرید از سرمایکس: %s
▫️فروش به سرمایکس: %s

%s
@sarmayex_finance
`

	//return fmt.Sprintf("قیمت تتر:\n\nخرید از سرمایکس: %s\nفروش به سرمایکس: %s \nپیج @sarmayex.finance", s1, s2)
	return fmt.Sprintf(msg, s1, s2, n)

}
