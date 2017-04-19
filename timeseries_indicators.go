package talib4g

import (
	"github.com/wcharczuk/go-chart"
	"github.com/wcharczuk/go-chart/drawing"
	"time"
)

type AmountIndictor struct {
	*TimeSeries
}

type pl func(i int) float64

func plotter(xvals []time.Time, ts []*Tick, datasource pl) chart.TimeSeries {
	y := make([]float64, len(ts))
	for i := range ts {
		y[i] = datasource(i)
	}
	return chart.TimeSeries{
		YValues: y,
		XValues: xvals,
		Style: chart.Style{
			Show:        true,
			StrokeWidth: 1,
			StrokeColor: drawing.ColorFromHex("3da9df"),
		},
	}
}

func (this AmountIndictor) Calculate(index int) float64 {
	return this.Ticks[index].Amount
}

func (this AmountIndictor) Plot(xvalues []time.Time) chart.TimeSeries {
	return plotter(xvalues, this.Ticks, func(i int) float64 {
		return this.Ticks[i].Amount
	})
}

type VolumeIndicator struct {
	*TimeSeries
}

func (this VolumeIndicator) Calculate(index int) float64 {
	return this.Ticks[index].Volume
}

type ClosePriceIndicator struct {
	*TimeSeries
}

func (this ClosePriceIndicator) Plot(xvals []time.Time) chart.TimeSeries {
	return plotter(xvals, this.Ticks, func(i int) float64 {
		return this.Ticks[i].ClosePrice
	})
}

func (this ClosePriceIndicator) Calculate(index int) float64 {
	return this.Ticks[index].ClosePrice
}

type TypicalPriceIndicator struct {
	*TimeSeries
}

func (this TypicalPriceIndicator) Calculate(index int) float64 {
	return (this.Ticks[index].MaxPrice + this.Ticks[index].MinPrice + this.Ticks[index].ClosePrice) / 3.0
}
