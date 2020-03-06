package app

import (
	"context"
	"github.com/adshao/go-binance"
	"strings"
)

const (
	errorValue = ""
)

func balanceToUSD(client *binance.Client, bal *binance.Balance) (float64, error) {
	haveFree := strToF(bal.Free)
	haveLocked := strToF(bal.Locked)
	if bal.Asset == "USDT" {
		return haveFree + haveLocked, nil
	}

	symbolPrice, err := client.NewListPricesService().Symbol(bal.Asset + "USDT").Do(context.Background())
	if err != nil {
		return 0, err
	}
	price := strToF(symbolPrice[0].Price)
	haveFree *= price
	haveLocked *= price
	return haveFree + haveLocked, nil
}

func binanceRateQuery(client *binance.Client) (string, error) {
	symbolPrice, err := client.NewListPricesService().Symbol("BTCUSDT").Do(context.Background())
	if err != nil {
		return "", err
	}
	return symbolPrice[0].Price, nil
}

func isEmptyBalance(str string) bool {
	return strings.Trim(str, ".0") == ""
}
