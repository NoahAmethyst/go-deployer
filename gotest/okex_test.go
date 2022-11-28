package gotest

import (
	"auto-deployer/utils/log"
	"github.com/tristan-club/wizard/pkg/http_util"
	"strconv"
	"testing"
)

func Test_Price(t *testing.T) {
	//sympols := []string{
	//	"ETHUSDT",
	//	"BTCUSDT",
	//	"LUNAUSDT",
	//}

	type CoinPrice struct {
		Symbol             string `json:"symbol"`
		LastPrice          string `json:"lastPrice"`
		PriceChangePercent string `json:"priceChangePercent"`
		OpenPrice          string `json:"openPrice"`
		HighPrice          string `json:"highPrice"`
		LowPrice           string `json:"lowPrice"`
	}

	data := CoinPrice{}
	err := http_util.GetJSON("https://api1.binance.com/api/v3/ticker/24hr?symbol=ETHUSDT", map[string]string{"Accepts": "application/json"}, &data)
	if err != nil {
		log.Error().Fields(map[string]interface{}{"action": "request coin pricee", "error": err.Error()}).Send()
		return
	}

	log.Info().Fields(map[string]interface{}{
		"info": data,
	}).Send()

	price, _ := strconv.ParseFloat(data.LastPrice, 64)
	log.Info().Msgf("%f", price)

}
