package formula

import "github.com/shopspring/decimal"

//	Интерфейс описывает функцию от одной переменной now с заданными парамтерами start и rate
// 	По смыслу start и rate задают время активации триггера и начальный курс
//	now - текущее время
//	Для start и now используется формат time.Unix
// 	Пример: rate-10+0.0002*(now-start)^1.2
//	now float64 - это костыль для построения графика, пока нет возможности с этим что-то сделать
type Formula interface {
	Calc(now float64) float64          // Возвращает значение в точке now
	CalcDec(now int64) decimal.Decimal // Возвращает значение в точке now в decimal.Decimal
	Start() int64                      // Возвращает значение параметра start
	Rate() decimal.Decimal             // Возвращает значение параметра rate
	String() string                    // Возвращает формулу в виде строки
	Alter(s string) error              // Заменяет формулу на новую, сохраняя старые rate и start
}
