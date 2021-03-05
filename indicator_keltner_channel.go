package techan

import (
	"github.com/sdcoffey/big"
)

type keltnerChannelIndicator struct {
	ema    Indicator
	atr    Indicator
	mul    big.Decimal
	window int
}

func NewKeltnerChannelUpperIndicator(series *TimeSeries, window int) Indicator {
	return keltnerChannelIndicator{
		atr:    NewAverageTrueRangeIndicator(series, window),
		ema:    NewEMAIndicator(NewClosePriceIndicator(series), window),
		mul:    big.ONE,
		window: window,
	}
}

func NewKeltnerChannelLowerIndicator(series *TimeSeries, window int) Indicator {
	return keltnerChannelIndicator{
		atr:    NewAverageTrueRangeIndicator(series, window),
		ema:    NewEMAIndicator(NewClosePriceIndicator(series), window),
		mul:    big.ONE.Neg(),
		window: window,
	}
}

func (kci keltnerChannelIndicator) Calculate(index int) big.Decimal {
	if index <= kci.window-1 {
		return big.ZERO
	}

	coefficient := big.NewFromInt(2).Mul(kci.mul)

	return kci.ema.Calculate(index).Add(kci.atr.Calculate(index).Mul(coefficient))
}
