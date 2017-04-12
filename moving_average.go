package talib4g

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
