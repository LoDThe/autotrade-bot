package formula

import (
	"time"

	"github.com/rwlist/autotrade-bot/pkg/convert"
	"github.com/shopspring/decimal"
)

//	Базовая функция удовлетворяющая интерфейсу Formula
//	Имеет вид rate-10+0.0002*(now-start)^1.2
//	Парсится через regexp patternBasic
type Basic struct {
	rate  decimal.Decimal
	start time.Time
	coef  []decimal.Decimal // Числовые коэффициенты
	orig  string
}

func (f *Basic) String() string {
	return f.orig
}

// Calc(now) Вычисляет значение в точке now
func (f *Basic) Calc(now time.Time) decimal.Decimal {
	t := now.Unix() - f.Start().Unix()
	brackets := decimal.NewFromInt(t)
	tmp := f.coef[1].Mul(convert.Pow(brackets, f.coef[2])) // Нужен Pow получше
	return f.Rate().
		Add(f.coef[0]).
		Add(tmp)
}

func (f *Basic) Start() time.Time {
	return f.start
}

func (f *Basic) Rate() decimal.Decimal {
	return f.rate
}

//	По заданной строке определяется является ли она формулой этого вида
//	Создаёт и возвращает указатель на структуру, если да
func NewBasic(s string, rate decimal.Decimal, start time.Time) (*Basic, error) {
	coef, err := parseBasic(s)
	if err != nil {
		return nil, err
	}
	return &Basic{
		rate:  rate,
		start: start,
		coef:  coef,
		orig:  s,
	}, nil
}

// Меняет формулу сохранняя значения rate и start прежними
func (f *Basic) Alter(s string) error {
	coef, err := parseBasic(s)
	if err != nil {
		return err
	}
	f.coef = coef
	return nil
}
