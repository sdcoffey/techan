package test

import (
	. "github.com/sdcoffey/talib4g"
	"github.com/stretchr/testify/assert"
	"github.com/wcharczuk/go-chart"
	"os"
	"testing"
)

func TestCumulativeGainsIndicator(t *testing.T) {
	ts := MockTimeSeries(1, 2, 3, 5, 8, 13)

	cumGains := CumulativeGainsIndicator{
		Indicator: ClosePriceIndicator{ts},
		TimeFrame: 10,
	}

	assert.EqualValues(t, 0, cumGains.Calculate(0))
	assert.EqualValues(t, 1, cumGains.Calculate(1))
	assert.EqualValues(t, 2, cumGains.Calculate(2))
	assert.EqualValues(t, 4, cumGains.Calculate(3))
	assert.EqualValues(t, 7, cumGains.Calculate(4))
	assert.EqualValues(t, 12, cumGains.Calculate(5))
}

func TestCumulativeLossesIndicator(t *testing.T) {
	ts := MockTimeSeries(13, 8, 5, 3, 2, 1)

	cumGains := CumulativeLossesIndicator{
		Indicator: ClosePriceIndicator{ts},
		TimeFrame: 10,
	}

	assert.EqualValues(t, 0, cumGains.Calculate(0))
	assert.EqualValues(t, -5, cumGains.Calculate(1))
	assert.EqualValues(t, -8, cumGains.Calculate(2))
	assert.EqualValues(t, -10, cumGains.Calculate(3))
	assert.EqualValues(t, -11, cumGains.Calculate(4))
	assert.EqualValues(t, -12, cumGains.Calculate(5))
}

func TestAverageIndicator(t *testing.T) {
	ts := MockTimeSeries(1, 2, 3, 5, 8, 13)

	avgGains := AverageIndicator{
		Indicator: CumulativeGainsIndicator{
			Indicator: ClosePriceIndicator{ts},
			TimeFrame: 10,
		},
		TimeFrame: 10,
	}

	assert.EqualValues(t, 0, avgGains.Calculate(0))
	assert.EqualValues(t, 1.0/2.0, avgGains.Calculate(1))
	assert.EqualValues(t, 2.0/3.0, avgGains.Calculate(2))
	assert.EqualValues(t, 4.0/4.0, avgGains.Calculate(3))
	assert.EqualValues(t, 7.0/5.0, avgGains.Calculate(4))
	assert.EqualValues(t, 12.0/6.0, avgGains.Calculate(5))
}

func TestChart(t *testing.T) {
	ts := RandomTimeSeries(1000)

	cpi := ClosePriceIndicator{ts}
	ema := NewEMAIndicator(cpi, 15)
	sma := SMAIndicator{cpi, 15}

	f, _ := os.OpenFile("data.png", os.O_CREATE|os.O_TRUNC|os.O_RDWR, os.FileMode(0755))
	defer f.Close()
	xvals := ts.TimeXValues()
	ch := chart.Chart{
		YAxis: chart.YAxis{
			Style: chart.StyleShow(),
		},
		Series: []chart.Series{
			cpi.Plot(xvals),
			ema.Plot(xvals),
			sma.Plot(xvals),
		},
	}
	err := ch.Render(chart.PNG, f)
	if err != nil {
		panic(err)
	}
}
