package techan

// Composition metadata keeps warm-up placeholders out of downstream windows
// and lets cache invalidation traverse all built-in indicator inputs.

func (gli gainLossIndicator) FirstValidIndex() int { return FirstValidIndex(gli.Indicator) + 1 }

func (gli gainLossIndicator) dependencies() []Indicator { return []Indicator{gli.Indicator} }

func (ci cumulativeIndicator) FirstValidIndex() int { return FirstValidIndex(ci.Indicator) + 1 }

func (ci cumulativeIndicator) dependencies() []Indicator { return []Indicator{ci.Indicator} }

func (pgi percentChangeIndicator) FirstValidIndex() int { return FirstValidIndex(pgi.Indicator) + 1 }

func (pgi percentChangeIndicator) dependencies() []Indicator { return []Indicator{pgi.Indicator} }

func (ai averageIndicator) FirstValidIndex() int { return FirstValidIndex(ai.Indicator) }

func (ai averageIndicator) dependencies() []Indicator { return []Indicator{ai.Indicator} }

func (di DerivativeIndicator) FirstValidIndex() int { return FirstValidIndex(di.Indicator) + 1 }

func (di DerivativeIndicator) dependencies() []Indicator { return []Indicator{di.Indicator} }

func (rs relativeStrengthIndicator) FirstValidIndex() int {
	return max(FirstValidIndex(rs.avgGain), FirstValidIndex(rs.avgLoss))
}

func (rs relativeStrengthIndicator) dependencies() []Indicator {
	return []Indicator{rs.avgGain, rs.avgLoss}
}

func (rsi relativeStrengthIndexIndicator) FirstValidIndex() int {
	return FirstValidIndex(rsi.rsIndicator)
}

func (rsi relativeStrengthIndexIndicator) dependencies() []Indicator {
	return []Indicator{rsi.rsIndicator}
}

func (mdi meanDeviationIndicator) FirstValidIndex() int { return FirstValidIndex(mdi.movingAverage) }

func (mdi meanDeviationIndicator) dependencies() []Indicator { return []Indicator{mdi.movingAverage} }

func (sdi windowedStandardDeviationIndicator) FirstValidIndex() int {
	return FirstValidIndex(sdi.movingAverage)
}

func (sdi windowedStandardDeviationIndicator) dependencies() []Indicator {
	return []Indicator{sdi.movingAverage}
}

func (vi varianceIndicator) FirstValidIndex() int { return FirstValidIndex(vi.Indicator) }

func (vi varianceIndicator) dependencies() []Indicator { return []Indicator{vi.Indicator} }

func (sdi standardDeviationIndicator) FirstValidIndex() int { return FirstValidIndex(sdi.indicator) }

func (sdi standardDeviationIndicator) dependencies() []Indicator { return []Indicator{sdi.indicator} }

func (bbi bbandIndicator) FirstValidIndex() int {
	return max(FirstValidIndex(bbi.ma), FirstValidIndex(bbi.stdev))
}

func (bbi bbandIndicator) dependencies() []Indicator { return []Indicator{bbi.ma, bbi.stdev} }

func (ai aroonIndicator) FirstValidIndex() int { return FirstValidIndex(ai.indicator) + ai.window - 1 }

func (ai aroonIndicator) dependencies() []Indicator { return []Indicator{ai.indicator} }

func (mvi maximumValueIndicator) FirstValidIndex() int { return FirstValidIndex(mvi.indicator) }

func (mvi maximumValueIndicator) dependencies() []Indicator { return []Indicator{mvi.indicator} }

func (mvi minimumValueIndicator) FirstValidIndex() int { return FirstValidIndex(mvi.indicator) }

func (mvi minimumValueIndicator) dependencies() []Indicator { return []Indicator{mvi.indicator} }

func (mdi maximumDrawdownIndicator) FirstValidIndex() int { return FirstValidIndex(mdi.indicator) }

func (mdi maximumDrawdownIndicator) dependencies() []Indicator { return []Indicator{mdi.indicator} }

func (tli trendLineIndicator) FirstValidIndex() int { return FirstValidIndex(tli.indicator) + 1 }

func (tli trendLineIndicator) dependencies() []Indicator { return []Indicator{tli.indicator} }

func (d dIndicator) FirstValidIndex() int { return FirstValidIndex(d.k) + d.window - 1 }

func (d dIndicator) dependencies() []Indicator { return []Indicator{d.k} }

func (rvii relativeVigorIndexIndicator) FirstValidIndex() int {
	return max(FirstValidIndex(rvii.numerator), FirstValidIndex(rvii.denominator)) + 3
}

func (rvii relativeVigorIndexIndicator) dependencies() []Indicator {
	return []Indicator{rvii.numerator, rvii.denominator}
}

func (rvsn relativeVigorIndexSignalLine) FirstValidIndex() int {
	return FirstValidIndex(rvsn.relativeVigorIndex) + 3
}

func (rvsn relativeVigorIndexSignalLine) dependencies() []Indicator {
	return []Indicator{rvsn.relativeVigorIndex}
}

func (kci keltnerChannelIndicator) FirstValidIndex() int {
	return max(FirstValidIndex(kci.ema), FirstValidIndex(kci.atr))
}

func (kci keltnerChannelIndicator) dependencies() []Indicator { return []Indicator{kci.ema, kci.atr} }

func (ccii commidityChannelIndexIndicator) FirstValidIndex() int { return ccii.window - 1 }

func (tri trueRangeIndicator) FirstValidIndex() int { return 1 }

func (atr averageTrueRangeIndicator) FirstValidIndex() int { return atr.window }
