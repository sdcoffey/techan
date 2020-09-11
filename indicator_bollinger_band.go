package techan

import "github.com/sdcoffey/big"

type windowedStandardDeviationIndicator struct {
	Indicator
	movingAverage Indicator
	window        int
}

func NewWindowedStandardDeviationIndicator(ind Indicator, window int) Indicator {
	return windowedStandardDeviationIndicator{
		Indicator:     ind,
		movingAverage: NewSimpleMovingAverage(ind, window),
		window:        window,
	}
}

// Calculate returns the windowed standard deviation of a base indicator
func (sdi windowedStandardDeviationIndicator) Calculate(index int) big.Decimal {
	avg := sdi.movingAverage.Calculate(index)
	variance := big.ZERO
	for i := Max(0, index-sdi.window+1); i <= index; i++ {
		pow := sdi.Indicator.Calculate(i).Sub(avg).Pow(2)
		variance = variance.Add(pow)
	}
	realwindow := Min(sdi.window, index+1)

	return variance.Div(big.NewDecimal(float64(realwindow))).Sqrt()
}

type bbandIndicator struct {
	ma     Indicator
	stdev  Indicator
	muladd big.Decimal
}

// Usually window=20, sigma=2
func NewBollingerUpperBandIndicator(indicator Indicator, window int, sigma float64) Indicator {
	return bbandIndicator{
		ma:     NewSimpleMovingAverage(indicator, window),
		stdev:  NewWindowedStandardDeviationIndicator(indicator, window),
		muladd: big.NewDecimal(sigma),
	}
}

func NewBollingerLowerBandIndicator(indicator Indicator, window int, sigma float64) Indicator {
	return bbandIndicator{
		ma:     NewSimpleMovingAverage(indicator, window),
		stdev:  NewWindowedStandardDeviationIndicator(indicator, window),
		muladd: big.NewDecimal(-sigma),
	}
}

func (bbi bbandIndicator) Calculate(index int) big.Decimal {
	return bbi.ma.Calculate(index).Add(bbi.stdev.Calculate(index).Mul(bbi.muladd))
}
