package techan

import (
	"math"

	"github.com/sdcoffey/big"
)

type moneyFlowIndexIndicator struct {
	mfIndicator Indicator
	oneHundred  big.Decimal
}

// NewMoneyFlowIndexIndicator returns a derivative Indicator which returns the money flow index of the base indicator
// in a given time frame. A more in-depth explanation of money flow index can be found here:
// https://www.investopedia.com/terms/m/mfi.asp
func NewMoneyFlowIndexIndicator(series *TimeSeries, timeframe int) Indicator {
	return moneyFlowIndexIndicator{
		mfIndicator: NewMoneyFlowRatioIndicator(series, timeframe),
		oneHundred:  big.NewFromString("100"),
	}
}

func (mfi moneyFlowIndexIndicator) Calculate(index int) big.Decimal {
	moneyFlowRatio := mfi.mfIndicator.Calculate(index)

	return mfi.oneHundred.Sub(mfi.oneHundred.Div(big.ONE.Add(moneyFlowRatio)))
}

type moneyFlowRatioIndicator struct {
	typicalPrice Indicator
	volume       Indicator
	window       int
}

// NewMoneyFlowRatioIndicator returns a derivative Indicator which returns the money flow ratio of the base indicator
// in a given time frame. Money flow ratio is the positive money flow divided by the negative money flow during the
// same time frame
func NewMoneyFlowRatioIndicator(series *TimeSeries, timeframe int) Indicator {
	return moneyFlowRatioIndicator{
		typicalPrice: NewTypicalPriceIndicator(series),
		volume:       NewVolumeIndicator(series),
		window:       timeframe,
	}
}

func (mfr moneyFlowRatioIndicator) Calculate(index int) big.Decimal {
	if index < mfr.window-1 {
		return big.ZERO
	}

	positiveMoneyFlow := big.ZERO
	negativeMoneyFlow := big.ZERO

	typicalPrice := mfr.typicalPrice.Calculate(index)

	for i := index; i > index-mfr.window && i > 0; i-- {
		prevTypicalPrice := mfr.typicalPrice.Calculate(i - 1)

		if typicalPrice.GT(prevTypicalPrice) {
			positiveMoneyFlow = positiveMoneyFlow.Add(typicalPrice.Mul(mfr.volume.Calculate(i)))
		} else if typicalPrice.LT(prevTypicalPrice) {
			negativeMoneyFlow = negativeMoneyFlow.Add(typicalPrice.Mul(mfr.volume.Calculate(i)))
		}

		typicalPrice = prevTypicalPrice
	}

	if negativeMoneyFlow.EQ(big.ZERO) {
		return big.NewDecimal(math.Inf(1))
	}

	return positiveMoneyFlow.Div(negativeMoneyFlow)
}
