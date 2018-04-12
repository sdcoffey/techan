package techan

import "github.com/sdcoffey/big"

type relativeVigorIndexIndicator struct {
	numerator   Indicator
	denominator Indicator
}

// NewRelativeVigorIndexIndicator returns an Indicator which returns the index of the relative vigor of the prices of
// a sercurity. Relative Vigor Index is simply the difference of the previous four days' close and open prices divided
// by the difference between the previous four days high and low prices. A more in-depth explanation of relative vigor
// index can be found here: https://www.fidelity.com/learning-center/trading-investing/technical-analysis/technical-indicator-guide/relative-vigor-index
func NewRelativeVigorIndexIndicator(series *TimeSeries) Indicator {
	return relativeVigorIndexIndicator{
		numerator:   NewDifferenceIndicator(NewClosePriceIndicator(series), NewOpenPriceIndicator(series)),
		denominator: NewDifferenceIndicator(NewHighPriceIndicator(series), NewLowPriceIndicator(series)),
	}
}

func (rvii relativeVigorIndexIndicator) Calculate(index int) big.Decimal {
	if index < 3 {
		return big.ZERO
	}

	two := big.NewFromString("2")

	a := rvii.numerator.Calculate(index)
	b := rvii.numerator.Calculate(index - 1).Mul(two)
	c := rvii.numerator.Calculate(index - 2).Mul(two)
	d := rvii.numerator.Calculate(index - 3)

	num := (a.Add(b).Add(c).Add(d)).Div(big.NewFromString("6"))

	e := rvii.denominator.Calculate(index)
	f := rvii.denominator.Calculate(index - 1).Mul(two)
	g := rvii.denominator.Calculate(index - 2).Mul(two)
	h := rvii.denominator.Calculate(index - 3)

	denom := (e.Add(f).Add(g).Add(h)).Div(big.NewFromString("6"))

	return num.Div(denom)
}

type relativeVigorIndexSignalLine struct {
	relativeVigorIndex Indicator
}

// NewRelativeVigorSignalLine returns an Indicator intended to be used in conjunction with Relative vigor index, which
// returns the average value of the last 4 indices of the RVI indicator.
func NewRelativeVigorSignalLine(series *TimeSeries) Indicator {
	return relativeVigorIndexSignalLine{
		relativeVigorIndex: NewRelativeVigorIndexIndicator(series),
	}
}

func (rvsn relativeVigorIndexSignalLine) Calculate(index int) big.Decimal {
	if index < 3 {
		return big.ZERO
	}

	rvi := rvsn.relativeVigorIndex.Calculate(index)
	i := rvsn.relativeVigorIndex.Calculate(index - 1).Mul(big.NewFromString("2"))
	j := rvsn.relativeVigorIndex.Calculate(index - 2).Mul(big.NewFromString("2"))
	k := rvsn.relativeVigorIndex.Calculate(index - 3)

	return (rvi.Add(i).Add(j).Add(k)).Div(big.NewFromString("6"))
}
