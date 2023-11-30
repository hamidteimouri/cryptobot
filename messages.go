package main

import (
	"context"
	"fmt"
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

	return fmt.Sprintf("قیمت تتر:\n\nخرید از سرمایکس: %s\nفروش به سرمایکس: %s", r.Sell.Round(0).String(), r.Buy.Round(0).String())

}
