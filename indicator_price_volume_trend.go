package techan

import "github.com/sdcoffey/big"

type pvtIndicator struct {
	closePriceChangeIndicator Indicator
	volumeIndicator           Indicator
	window                    int
	resultCache               resultCache
}

// NewPriceVolumeTrendIndicator is a derivative indicator that returns the Price Volume Trend (also known as
// Volume Price Trend or VPT) for a given window. A more in-depth explanation can be found here:
// https://www.investopedia.com/terms/v/vptindicator.asp
func NewPriceVolumeTrendIndicator(closePriceIndicator, volumeIndicator Indicator, window int) Indicator {
	return &pvtIndicator{
		closePriceChangeIndicator: NewPercentChangeIndicator(closePriceIndicator),
		volumeIndicator:           volumeIndicator,
		window:                    window,
		resultCache:               make([]*big.Decimal, 1000),
	}
}

func (pvt *pvtIndicator) Calculate(index int) big.Decimal {
	if cachedValue := returnIfCached(pvt, index, func(i int) big.Decimal {
		return big.ZERO
	}); cachedValue != nil {
		return *cachedValue
	}

	priceVolumeChange := pvt.volumeIndicator.Calculate(index).Mul(pvt.closePriceChangeIndicator.Calculate(index))
	previousPVT := pvt.Calculate(index - 1)
	result := priceVolumeChange.Add(previousPVT)

	cacheResult(pvt, index, result)

	return result
}

func (pvt *pvtIndicator) cache() resultCache { return pvt.resultCache }

func (pvt *pvtIndicator) setCache(newCache resultCache) {
	pvt.resultCache = newCache
}

func (pvt *pvtIndicator) windowSize() int { return pvt.window }

// NewPVTAndSignalIndicator returns the difference between the Price Volume Trend and a signal line.
// A signal line can be and EMA or SMA of the Price Volume Trend
func NewPVTAndSignalIndicator(pvtIndicator, signalIndicator Indicator) Indicator {
	return NewDifferenceIndicator(pvtIndicator, signalIndicator)
}
