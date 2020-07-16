package logic

import (
	"fmt"

	"github.com/rwlist/autotrade-bot/pkg/trade"

	"github.com/rwlist/autotrade-bot/pkg/trigger"
)

func startMessage(order *trade.Order) string {
	return fmt.Sprintf("A %v BTC/USDT order was placed with price = %v.\nWaiting for %s", order.Side, order.Price, sleepDur)
}

func orderStatusMessage(order *trade.Order) string {
	return fmt.Sprintf("Side: %v\nDone %v / %v\nStatus: %v", order.Side, order.ExecutedQuantity, order.OrigQuantity, order.Status)
}

func triggerResponseMessage(resp *trigger.Response) string {
	return fmt.Sprintf("Current rate: %v\nFormula rate: %.2f\n\n"+
		"Absolute difference: %.2f\nRelative difference: %.2f%%\n\n"+
		"Start rate: %v\nRelative profit: %.2f%%\nAbsolute profit: %.2f\n\n"+
		"Error: %v\nUpdate time: %v",
		resp.CurRate, resp.FormulaRate, resp.AbsDif, resp.RelDif, resp.StartRate, resp.RelProfit, resp.AbsProfit,
		resp.Err, resp.T.Format("02.01.2006 15.04.05"))
}
