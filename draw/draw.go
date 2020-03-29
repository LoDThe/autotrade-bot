package draw

import (
	"fmt"
	"image/color"
	"time"

	"github.com/rwlist/autotrade-bot/to_str"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
)

type AllTimeTicks struct{}

func (AllTimeTicks) Ticks(min, max float64) []plot.Tick {
	tks := plot.DefaultTicks{}.Ticks(min, max)
	for i, t := range tks {
		ut := time.Unix(int64(t.Value), 0)
		d := to_str.Str(ut.Day())
		h := to_str.Str(ut.Hour())
		m := to_str.Str(ut.Minute())
		if len(d) == 1 {
			d = "0" + d
		}
		if len(h) == 1 {
			h = "0" + h
		}
		if len(m) == 1 {
			m = "0" + m
		}
		tks[i].Label = fmt.Sprintf("%v.%v\n%v:%v", d, ut.Month(), h, m)
	}
	return tks
}

type AllPriceTicks struct{}

func (AllPriceTicks) Ticks(min, max float64) []plot.Tick {
	tks := plot.DefaultTicks{}.Ticks(min, max)
	for i, t := range tks {
		if t.Label == "" {
			tks[i].Label = to_str.Str64(int64(t.Value))
		}
	}
	return tks
}

func MakeHorLine(x, y float64, r, g, b uint8) *plotter.Line {
	pts := make(plotter.XYs, 2)
	pts[0].Y = y
	pts[1].Y = y
	pts[0].X = x - 1000
	pts[1].X = x + 86400 + 10000
	line, _ := plotter.NewLine(pts)
	line.Color = color.RGBA{
		R: r,
		G: g,
		B: b,
		A: 255,
	}
	return line
}
