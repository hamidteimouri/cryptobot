package main

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/hamidteimouri/gommon/htenvier"
	"github.com/shopspring/decimal"
	httpClientGoLib "github.com/trustwallet/go-libs/client"
)

type TetherPriceResponse struct {
	Cost      string `json:"cost"`
	SellPrice string `json:"sellPrice"`
	BuyPrice  string `json:"buyPrice"`
}

type TetherPrice struct {
	Buy    decimal.Decimal
	Sell   decimal.Decimal
	Avg    decimal.Decimal
	Supply decimal.Decimal
}

const (
	sarmayexGetTetherPrice = "api/v1/usdt-price"
)

func GetTetherPrice(ctx context.Context) (*TetherPrice, error) {

	req := httpClientGoLib.InitJSONClient(htenvier.Env("SARMAYEX_BASE_URL"), nil)
	nb := httpClientGoLib.NewReqBuilder()
	nb.Method("GET").PathStatic(sarmayexGetTetherPrice)

	body, err := req.Execute(ctx, nb.Build())
	if err != nil {
		return nil, err
	}

	var result TetherPriceResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	if result.BuyPrice == "" || result.SellPrice == "" {
		return nil, errors.New("tether prices are null from Sarmayex")
	}

	if result.Cost == "" {
		return nil, errors.New("tether supply price is null from Sarmayex")
	}

	sell, err := decimal.NewFromString(result.SellPrice)
	if err != nil {
		return nil, err
	}

	buy, err := decimal.NewFromString(result.BuyPrice)
	if err != nil {
		return nil, err
	}

	cost, err := decimal.NewFromString(result.Cost)
	if err != nil {
		return nil, err
	}

	two := decimal.New(2, 0)
	r := &TetherPrice{
		Buy:    buy,
		Sell:   sell,
		Avg:    buy.Add(sell).Div(two).RoundDown(0),
		Supply: cost.RoundDown(0), // this is supply price of tether of Sarmayex
	}

	return r, nil
}
