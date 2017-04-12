package talib4g

type RSIIndicator struct {
	AvgGainIndicator AverageIndicator
	AvgLossIndicator AverageIndicator
	TimeFrame        int
}

func NewRSIIndicator(ind Indicator, timeFrame int) RSIIndicator {
	return RSIIndicator{
		AvgGainIndicator: AverageIndicator{CumulativeGainsIndicator{ind, timeFrame}, timeFrame},
		AvgLossIndicator: AverageIndicator{CumulativeLossesIndicator{ind, timeFrame}, timeFrame},
		TimeFrame:        timeFrame,
	}
}

func (this RSIIndicator) Calculate(index int) float64 {
	if index == 0 {
		return 0.0
	}

	averageGain := this.AvgGainIndicator.Calculate(index)
	averageLoss := this.AvgLossIndicator.Calculate(index)

	relativeStrength := 0.0
	if averageLoss > 0 {
		relativeStrength = averageGain / averageLoss
	}

	return 100.0 - (100.0 / (1 + relativeStrength))
}
