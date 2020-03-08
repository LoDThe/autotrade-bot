package main

import (
	"encoding/json"
	"github.com/adshao/go-binance"
	"github.com/petuhovskiy/telegram"
	"github.com/petuhovskiy/telegram/updates"
	"log"

	"github.com/petuhovskiy/autotrade-bot/app"
	"github.com/petuhovskiy/autotrade-bot/conf"
)

func main() {
	cfg, err := conf.ParseEnv()
	if err != nil {
		log.Fatal(err)
	}

	bot := telegram.NewBotWithOpts(cfg.Bot.Token, &telegram.Opts{
		Middleware: func(handler telegram.RequestHandler) telegram.RequestHandler {
			return func(methodName string, req interface{}) (message json.RawMessage, err error) {
				res, err := handler(methodName, req)
				if err != nil {
					log.Println("Telegram response error: ", err)
				}

				return res, err
			}
		},
	})

	ch, err := updates.StartPolling(bot, telegram.GetUpdatesRequest{
		Offset:         0,
		Limit:          50,
		Timeout:        10,
	})
	if err != nil {
		log.Fatal(err)
	}

	c := binance.NewClient(cfg.Binance.APIKey, cfg.Binance.Secret)
	l := app.NewLogic(c)
	h := app.NewHandler(bot, l, cfg)

	for upd := range ch {
		h.Handle(upd)
	}
}