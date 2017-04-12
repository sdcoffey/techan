package talib4g

type CCIIndicator struct {
	tpi TypicalPriceIndicator
	sma SMAIndicator
	md  MeanDeviationIndicator
	tf  int
}

func NewCCIIndicator(ts *TimeSeries, timeFrame int) CCIIndicator {
	return CCIIndicator{
		tpi: TypicalPriceIndicator{ts},
		sma: SMAIndicator{TypicalPriceIndicator{ts}, timeFrame},
		md:  NewMeanDeviationIndicator(TypicalPriceIndicator{ts}, timeFrame),
		tf:  timeFrame,
	}
}

func (this CCIIndicator) Calculate(index int) float64 {
	typicalPrice := this.tpi.Calculate(index)
	typicalPriceAvg := this.sma.Calculate(index)
	mean := this.md.Calculate(index)

	return (typicalPrice - typicalPriceAvg) / (mean * 0.015)
}
