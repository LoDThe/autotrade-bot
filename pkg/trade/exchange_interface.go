package trade

import "github.com/rwlist/autotrade-bot/pkg/draw"

type IExchange interface {
	AccountBalance() ([]Balance, error)
	AccountSymbolBalance(symbol string) (float64, error)
	BalanceToUSD(bal *Balance) (float64, error)
	GetRate(symbol ...string) (string, error)
	BuyAll(symbol ...string) *Status
	SellAll(symbol ...string) *Status
	GetOrder(id int64, symbol ...string) (*Order, error)
	CancelOrder(id int64, symbol ...string) error
	GetKlines(opts ...draw.KlinesOpts) (*draw.Klines, error)
	SetScale(scale string)
	SetSymbol(symbol string)
}
