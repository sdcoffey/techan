package talib4g

import (
	"github.com/wcharczuk/go-chart"
	"github.com/wcharczuk/go-chart/drawing"
	"time"
)

type SMAIndicator struct {
	Indicator Indicator
	TimeFrame int
}

func (this SMAIndicator) Calculate(index int) float64 {
	sum := 0.0
	for i := Max(0, index-this.TimeFrame+1); i <= index; i++ {
		sum += this.Indicator.Calculate(i)
	}
	realTimeFrame := Min(this.TimeFrame, index+1)

	return sum / float64(realTimeFrame)
}

func (this SMAIndicator) Plot(xvals []time.Time) chart.TimeSeries {
	y := make([]float64, len(xvals))
	for i := range xvals {
		y[i] = this.Calculate(i)
	}
	return chart.TimeSeries{
		YValues: y,
		XValues: xvals,
		Style: chart.Style{
			Show:        true,
			StrokeWidth: 1,
			StrokeColor: drawing.ColorFromHex("f00"),
		},
	}
}

type EMAIndicator struct {
	Indicator   Indicator
	TimeFrame   int
	resultCache []float64
}

func NewEMAIndicator(indicator Indicator, timeFrame int) *EMAIndicator {
	return &EMAIndicator{
		Indicator:   indicator,
		TimeFrame:   timeFrame,
		resultCache: make([]float64, timeFrame),
	}
}

func (this *EMAIndicator) Calculate(index int) float64 {
	if index+1 < this.TimeFrame {
		return SMAIndicator{this.Indicator, this.TimeFrame}.Calculate(index)
	}

	if index == 0 {
		result := this.Indicator.Calculate(index)
		return result
	}

	emaPrev := this.Calculate(index - 1)
	mult := 2.0 / float64(this.TimeFrame+1)
	result := (this.Indicator.Calculate(index)-emaPrev)*mult + emaPrev

	return result
}

func (this *EMAIndicator) Plot(xvals []time.Time) chart.TimeSeries {
	y := make([]float64, len(xvals))
	for i := range xvals {
		y[i] = this.Calculate(i)
	}
	return chart.TimeSeries{
		YValues: y,
		XValues: xvals,
		Style: chart.Style{
			Show:        true,
			StrokeWidth: 1,
			StrokeColor: drawing.ColorFromHex("0f0"),
		},
	}
}

func (this *EMAIndicator) cacheResult(index int, val float64) {
	if index < len(this.resultCache) {
		this.resultCache[index] = val
	} else {
		this.resultCache = append(this.resultCache, val)
	}
}

func (this EMAIndicator) multiplier(index int) float64 {
	return 2.0 / (float64(index) + 1)
}

type MACDIndicator struct {
	shortEMA Indicator
	longEMA  Indicator
}

func NewMACDIndicator(i Indicator, shortTimeFrame, longTimeFrame int) MACDIndicator {
	return MACDIndicator{
		NewEMAIndicator(i, shortTimeFrame),
		NewEMAIndicator(i, longTimeFrame),
	}
}

func (this MACDIndicator) Calculate(index int) float64 {
	return this.shortEMA.Calculate(index) - this.longEMA.Calculate(index)
}
