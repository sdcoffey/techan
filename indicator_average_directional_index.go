package techan

import (
	"github.com/sdcoffey/big"
)

type averageDirectionalIndexIndicator struct {
	series  *TimeSeries
	window  int
	plusDI  Indicator
	minusDI Indicator
}

// NewAverageDirectionalIndexIndicator returns an indicator that calculates the ADX
// which measures trend strength irrespective of direction
// window: number of periods (typically 14)
func NewAverageDirectionalIndexIndicator(series *TimeSeries, window int) Indicator {
	tr := NewTrueRangeIndicator(series)
	plusDM := NewDirectionalMovementIndicator(series, true)
	minusDM := NewDirectionalMovementIndicator(series, false)

	// Smooth TR, +DM, -DM using Wilder's smoothing
	smoothedTR := NewWilderSmoothing(tr, window)
	smoothedPlusDM := NewWilderSmoothing(plusDM, window)
	smoothedMinusDM := NewWilderSmoothing(minusDM, window)

	// Calculate +DI and -DI
	plusDI := NewDirectionalIndicator(smoothedPlusDM, smoothedTR)
	minusDI := NewDirectionalIndicator(smoothedMinusDM, smoothedTR)

	return &averageDirectionalIndexIndicator{
		series:  series,
		window:  window,
		plusDI:  plusDI,
		minusDI: minusDI,
	}
}

type directionalMovementIndicator struct {
	series *TimeSeries
	isPlus bool
}

func NewDirectionalMovementIndicator(series *TimeSeries, isPlus bool) Indicator {
	return &directionalMovementIndicator{
		series: series,
		isPlus: isPlus,
	}
}

func (dm *directionalMovementIndicator) Calculate(index int) big.Decimal {
	if index <= 0 {
		return big.ZERO
	}

	curr := dm.series.Candles[index]
	prev := dm.series.Candles[index-1]

	if dm.isPlus {
		// +DM = Current High - Previous High (if positive and greater than -DM)
		upMove := curr.MaxPrice.Sub(prev.MaxPrice)
		downMove := prev.MinPrice.Sub(curr.MinPrice)
		if upMove.GT(downMove) && upMove.GT(big.ZERO) {
			return upMove
		}
	} else {
		// -DM = Previous Low - Current Low (if positive and greater than +DM)
		upMove := curr.MaxPrice.Sub(prev.MaxPrice)
		downMove := prev.MinPrice.Sub(curr.MinPrice)
		if downMove.GT(upMove) && downMove.GT(big.ZERO) {
			return downMove
		}
	}
	return big.ZERO
}

type directionalIndicator struct {
	dm Indicator
	tr Indicator
}

func NewDirectionalIndicator(dm, tr Indicator) Indicator {
	return &directionalIndicator{
		dm: dm,
		tr: tr,
	}
}

func (di *directionalIndicator) Calculate(index int) big.Decimal {
	tr := di.tr.Calculate(index)
	if tr.EQ(big.ZERO) {
		return big.ZERO
	}

	// DI = (Smoothed DM / Smoothed TR) * 100
	return di.dm.Calculate(index).Div(tr).Mul(big.NewDecimal(100))
}

type wilderSmoothingIndicator struct {
	indicator Indicator
	window    int
}

func NewWilderSmoothing(indicator Indicator, window int) Indicator {
	return &wilderSmoothingIndicator{
		indicator: indicator,
		window:    window,
	}
}

func (ws *wilderSmoothingIndicator) Calculate(index int) big.Decimal {
	if index < ws.window {
		return ws.indicator.Calculate(index)
	}

	if index == ws.window {
		// First value is simple average
		sum := big.ZERO
		for i := 0; i < ws.window; i++ {
			sum = sum.Add(ws.indicator.Calculate(index - i))
		}
		return sum.Div(big.NewDecimal(float64(ws.window)))
	}

	// Subsequent values use Wilder's smoothing formula:
	// (Prior Value * (period - 1) + Current Value) / period
	prevValue := ws.Calculate(index - 1)
	currentValue := ws.indicator.Calculate(index)
	multiplier := big.NewDecimal(float64(ws.window - 1))

	return prevValue.Mul(multiplier).Add(currentValue).Div(big.NewDecimal(float64(ws.window)))
}

func (adx *averageDirectionalIndexIndicator) Calculate(index int) big.Decimal {
	if index < adx.window {
		return big.ZERO
	}

	plusDI := adx.plusDI.Calculate(index)
	minusDI := adx.minusDI.Calculate(index)

	// Calculate DX = |(+DI - -DI)| / (+DI + -DI) * 100
	sumDI := plusDI.Add(minusDI)
	if sumDI.EQ(big.ZERO) {
		return big.ZERO
	}

	diffDI := plusDI.Sub(minusDI).Abs()
	dx := diffDI.Div(sumDI).Mul(big.NewDecimal(100))

	// For first window periods, return DX
	if index == adx.window {
		return dx
	}

	// After window periods, use Wilder's smoothing for ADX
	prevADX := adx.Calculate(index - 1)
	return prevADX.Mul(big.NewDecimal(float64(adx.window - 1))).Add(dx).Div(big.NewDecimal(float64(adx.window)))
}
