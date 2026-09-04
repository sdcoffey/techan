package techan

import (
	"bytes"
	"math"
	"strings"
	"testing"
	"time"

	"github.com/sdcoffey/big"
)

func auditSeries(values ...float64) *TimeSeries {
	s := NewTimeSeries()
	for i, v := range values {
		c := NewCandle(NewTimePeriod(time.Unix(int64(i), 0), time.Second))
		c.OpenPrice, c.ClosePrice = big.NewDecimal(v), big.NewDecimal(v)
		c.MaxPrice, c.MinPrice = big.NewDecimal(v+1), big.NewDecimal(v-1)
		s.AddCandle(c)
	}
	return s
}

func auditEqual(t *testing.T, got big.Decimal, want float64) {
	t.Helper()
	if math.IsNaN(got.Float()) || math.Abs(got.Float()-want) > 1e-8 {
		t.Errorf("got %v, want %v", got.Float(), want)
	}
}

func TestRegressionCrossovers(t *testing.T) {
	t.Run("up must stop firing after crossing", func(t *testing.T) {
		r := NewCrossUpIndicatorRule(NewConstantIndicator(2), NewFixedIndicator(1, 3, 4))
		if r.IsSatisfied(2, nil) {
			t.Error("still true one bar after crossing")
		}
	})
	t.Run("down must detect the downward crossing", func(t *testing.T) {
		r := NewCrossDownIndicatorRule(NewFixedIndicator(3, 1), NewConstantIndicator(2))
		if !r.IsSatisfied(1, nil) {
			t.Error("false at downward crossing")
		}
	})
	t.Run("equal values have not crossed", func(t *testing.T) {
		r := NewCrossUpIndicatorRule(NewConstantIndicator(2), NewFixedIndicator(2, 2))
		if r.IsSatisfied(1, nil) {
			t.Error("equal values treated as a crossing")
		}
	})
}

func TestRegressionCompositeCacheReset(t *testing.T) {
	for name, makeIndicator := range map[string]func(*TimeSeries) Indicator{
		"MACD": func(s *TimeSeries) Indicator { return NewMACDIndicator(NewClosePriceIndicator(s), 2, 3) },
		"RSI":  func(s *TimeSeries) Indicator { return NewRelativeStrengthIndexIndicator(NewClosePriceIndicator(s), 3) },
		"nested EMA": func(s *TimeSeries) Indicator {
			return NewEMAIndicator(NewEMAIndicator(NewClosePriceIndicator(s), 2), 2)
		},
	} {
		t.Run(name, func(t *testing.T) {
			s := auditSeries(10, 12, 9, 11, 10, 12)
			ind := makeIndicator(s)
			before := ind.Calculate(5)
			s.Candles[5].ClosePrice = big.NewFromInt(30)
			reset := ResetCacheFrom(ind, 5)
			got, want := ind.Calculate(5), makeIndicator(s).Calculate(5)
			if got.Sub(want).Abs().GT(big.NewDecimal(1e-8)) {
				t.Errorf("reset=%v before=%v after=%v fresh=%v", reset, before.Float(), got.Float(), want.Float())
			}
		})
	}
}

func TestRegressionMaximumDrawdown(t *testing.T) {
	auditEqual(t, NewMaximumDrawdownIndicator(NewFixedIndicator(100, 200, 300), -1).Calculate(2), 0)
}

func TestRegressionCCI(t *testing.T) {
	s := auditSeries(10, 11, 12)
	s.Candles[2].MaxPrice = big.NewFromInt(24)
	s.Candles[2].MinPrice = big.NewFromInt(12)
	// Typical prices are 10, 11, 16. Mean=37/3; mean absolute deviation=22/9.
	auditEqual(t, NewCCIIndicator(s, 3).Calculate(2), 100)
}

func TestRegressionMACDWarmupContaminatesSignal(t *testing.T) {
	s := auditSeries(100, 100, 100, 100, 100, 100, 100, 100)
	macd := NewMACDIndicator(NewClosePriceIndicator(s), 2, 4)
	auditEqual(t, NewMACDHistogramIndicator(macd, 2).Calculate(4), 0)
}

func TestRegressionBollingerWarmup(t *testing.T) {
	// A constant series has zero standard deviation, including partial windows.
	auditEqual(t, NewWindowedStandardDeviationIndicator(NewFixedIndicator(10, 10, 10), 3).Calculate(0), 0)
}

func TestRegressionAroonOrderIndependence(t *testing.T) {
	base := NewFixedIndicator(1, 5, 2, 3, 4)
	ind := NewAroonUpIndicator(base, 3)
	ind.Calculate(2)
	got := ind.Calculate(4)
	want := NewAroonUpIndicator(base, 3).Calculate(4)
	if !got.EQ(want) {
		t.Errorf("skipped indices: got %v, fresh %v", got.Float(), want.Float())
	}
	ind = NewAroonUpIndicator(base, 3)
	ind.Calculate(4)
	got = ind.Calculate(2)
	want = NewAroonUpIndicator(base, 3).Calculate(2)
	if !got.EQ(want) {
		t.Errorf("backwards indices: got %v, fresh %v", got.Float(), want.Float())
	}
}

func auditRecord(side OrderSide, entry, exit float64, duration time.Duration) *TradingRecord {
	r := NewTradingRecord()
	r.Operate(Order{Side: side, Security: "XYZ", Amount: big.ONE, Price: big.NewDecimal(entry), ExecutionTime: time.Unix(0, 0)})
	r.Operate(Order{Side: 1 - side, Security: "XYZ", Amount: big.ONE, Price: big.NewDecimal(exit), ExecutionTime: time.Unix(0, 0).Add(duration)})
	return r
}

func TestRegressionShortAnalysis(t *testing.T) {
	r := auditRecord(SELL, 100, 90, time.Hour)
	if got := (PercentGainAnalysis{}).Analyze(r); math.Abs(got-0.1) > 1e-8 {
		t.Errorf("percent gain: got %v, want +0.1", got)
	}
	var buf bytes.Buffer
	LogTradesAnalysis{Writer: &buf}.Analyze(r)
	if !strings.Contains(buf.String(), "Profit: $10") {
		t.Errorf("short log: %s", buf.String())
	}
}

func TestRegressionPeriodProfit(t *testing.T) {
	r := auditRecord(BUY, 100, 115, 90*time.Minute)
	if got := (PeriodProfitAnalysis{Period: time.Hour}).Analyze(r); got != 10 {
		t.Errorf("got %v per hour, want 10", got)
	}
}

func TestRegressionBuyAndHoldHorizon(t *testing.T) {
	s := auditSeries(100, 110, 200)
	s.Candles[2].Period = NewTimePeriod(time.Unix(3, 0), time.Second)
	r := auditRecord(BUY, 100, 110, 2*time.Second)
	got := (BuyAndHoldAnalysis{TimeSeries: s, StartingMoney: 100}).Analyze(r)
	if got != 10 {
		t.Errorf("record ends at t=2; got %v using t=3..4 price, want 10", got)
	}
}

func TestRegressionParseTimePeriod(t *testing.T) {
	t.Run("invalid text must return error", func(t *testing.T) {
		if got, err := ParseTimePeriod("garbage"); err == nil {
			t.Errorf("accepted garbage: %v", got)
		}
	})
	t.Run("three dates must not panic", func(t *testing.T) {
		defer func() {
			if v := recover(); v != nil {
				t.Errorf("panicked: %v", v)
			}
		}()
		ParseTimePeriod("2020-01-01 -> 2020-01-02 -> 2020-01-03")
	})
}

func TestRegressionRSISeed(t *testing.T) {
	// Three changes: +1, -1, +1 => average gain=2/3 and average loss=1/3 => RSI=200/3.
	auditEqual(t, NewRelativeStrengthIndexIndicator(NewFixedIndicator(1, 2, 1, 2), 3).Calculate(3), 200.0/3)
}

func TestRegressionWarmupBoundary(t *testing.T) {
	r := OverIndicatorRule{First: NewConstantIndicator(2), Second: NewConstantIndicator(1)}
	s := RuleStrategy{EntryRule: r, ExitRule: r, UnstablePeriod: 0}
	if !s.ShouldEnter(0, NewTradingRecord()) {
		t.Error("zero warmup still skips bar 0")
	}
}

func TestRegressionCompositeMutationAcrossGraph(t *testing.T) {
	factories := map[string]func(Indicator) Indicator{
		"SMA":                func(i Indicator) Indicator { return NewSimpleMovingAverage(i, 3) },
		"MMA":                func(i Indicator) Indicator { return NewMMAIndicator(i, 3) },
		"MACD histogram":     func(i Indicator) Indicator { return NewMACDHistogramIndicator(NewMACDIndicator(i, 2, 4), 3) },
		"Bollinger":          func(i Indicator) Indicator { return NewBollingerUpperBandIndicator(i, 3, 2) },
		"deviation":          func(i Indicator) Indicator { return NewMeanDeviationIndicator(i, 3) },
		"variance":           NewVarianceIndicator,
		"standard deviation": NewStandardDeviationIndicator,
		"derivative":         func(i Indicator) Indicator { return DerivativeIndicator{Indicator: i} },
		"RSI":                func(i Indicator) Indicator { return NewRelativeStrengthIndexIndicator(i, 3) },
		"average gains":      func(i Indicator) Indicator { return NewAverageGainsIndicator(i, 3) },
		"percent change":     NewPercentChangeIndicator,
		"Aroon":              func(i Indicator) Indicator { return NewAroonUpIndicator(i, 3) },
		"drawdown":           func(i Indicator) Indicator { return NewMaximumDrawdownIndicator(i, 3) },
		"maximum":            func(i Indicator) Indicator { return NewMaximumValueIndicator(i, 3) },
		"minimum":            func(i Indicator) Indicator { return NewMinimumValueIndicator(i, 3) },
		"trend":              func(i Indicator) Indicator { return NewTrendlineIndicator(i, 3) },
		"slow stochastic":    func(i Indicator) Indicator { return NewSlowStochasticIndicator(i, 3) },
	}
	for name, factory := range factories {
		t.Run(name, func(t *testing.T) {
			s := auditSeries(10, 12, 9, 11, 10, 12, 14, 13, 15, 12)
			makeIndicator := func() Indicator { return factory(NewEMAIndicator(NewClosePriceIndicator(s), 2)) }
			actual := makeIndicator()
			actual.Calculate(9)
			s.Candles[3].ClosePrice = big.NewFromInt(30)
			if !ResetCacheFrom(actual, 3) {
				t.Fatal("did not reach cached input")
			}
			expected := makeIndicator()
			for i := 3; i < 10; i++ {
				auditEqual(t, actual.Calculate(i), expected.Calculate(i).Float())
			}
		})
	}
}

func TestRegressionWarmupMetadata(t *testing.T) {
	price := NewFixedIndicator(10, 11, 12, 13, 14, 15, 16, 17, 18, 19)
	macd := NewMACDIndicator(price, 2, 4)
	cases := []struct {
		indicator Indicator
		first     int
	}{
		{NewEMAIndicator(price, 3), 2},
		{NewSimpleMovingAverage(NewEMAIndicator(price, 3), 4), 5},
		{NewEMAIndicator(NewEMAIndicator(price, 3), 4), 5},
		{macd, 3},
		{NewMACDHistogramIndicator(macd, 3), 5},
		{NewRelativeStrengthIndexIndicator(price, 3), 3},
	}
	for _, tc := range cases {
		if got := FirstValidIndex(tc.indicator); got != tc.first {
			t.Errorf("first index=%d want=%d", got, tc.first)
		}
		for i := 0; i < tc.first; i++ {
			auditEqual(t, tc.indicator.Calculate(i), 0)
		}
	}
	// Start the signal from the first three valid MACD observations.
	want := macd.Calculate(3).Add(macd.Calculate(4)).Add(macd.Calculate(5)).Div(big.NewFromInt(3))
	got := NewEMAIndicator(macd, 3).Calculate(5)
	auditEqual(t, got, want.Float())
}

func TestRegressionSmoothingPrecisionAndScale(t *testing.T) {
	for name, factory := range map[string]func(Indicator, int) Indicator{"EMA": NewEMAIndicator, "MMA": NewMMAIndicator} {
		t.Run(name, func(t *testing.T) {
			values := make([]float64, 1600)
			for i := range values {
				values[i] = float64(100 + i%7)
			}
			ind := factory(NewFixedIndicator(values...), 10)
			last := ind.Calculate(len(values) - 1)
			if digits := len(last.FormattedString(-1)); digits > 100 {
				t.Fatalf("recursive state grew to %d digits", digits)
			}
			ResetCacheFrom(ind, 500)
			auditEqual(t, ind.Calculate(len(values)-1), last.Float())
			for _, value := range []string{"0.123456789012345678901234567890123456789", "1e-400", "1e400"} {
				s := mockTimeSeries(value, value, value, value, value)
				got := factory(NewClosePriceIndicator(s), 3).Calculate(4)
				want := big.NewFromString(value)
				relativeError := got.Sub(want).Abs().Div(want.Abs())
				if relativeError.GT(big.NewFromString("1e-70")) {
					t.Errorf("lost precision for %s: %s", value, relativeError)
				}
			}
		})
	}
}

func TestRegressionFlatAndMonotoneIndicators(t *testing.T) {
	for _, values := range [][]float64{{100, 100, 100, 100, 100, 100}, {100, 200, 300, 400, 500, 600}} {
		drawdown := NewMaximumDrawdownIndicator(NewFixedIndicator(values...), 3)
		for i := range values {
			auditEqual(t, drawdown.Calculate(i), 0)
		}
	}
	s := auditSeries(10, 10, 10, 10, 10)
	for i := 0; i < 5; i++ {
		auditEqual(t, NewCCIIndicator(s, 3).Calculate(i), 0)
		auditEqual(t, NewWindowedStandardDeviationIndicator(NewClosePriceIndicator(s), 3).Calculate(i), 0)
		if i < 2 {
			auditEqual(t, NewBollingerLowerBandIndicator(NewClosePriceIndicator(s), 3, 2).Calculate(i), 0)
		}
	}
}

func TestRegressionCrossoverPlateaus(t *testing.T) {
	for _, factory := range []func(Indicator, Indicator) Rule{NewCrossUpIndicatorRule, NewCrossDownIndicatorRule} {
		// Both constructors take the initially upper indicator first.
		upper, lower := NewFixedIndicator(3, 2, 2, 1, 0), NewFixedIndicator(1, 2, 2, 3, 4)
		rule := factory(upper, lower)
		for i, want := range []bool{false, false, false, true, false} {
			if got := rule.IsSatisfied(i, nil); got != want {
				t.Errorf("index %d got %v want %v", i, got, want)
			}
		}
	}
}

func TestRegressionParserValidation(t *testing.T) {
	for _, s := range []string{"", "garbage", "2020-01-01 garbage", "garbage 2020-01-01", "2020-01-01T12:00", "2020-13-01", "2020-01-01 2020-01-02 2020-01-03", "2020-01-02 -> 2020-01-01", "2020-01-01T12:00:00Z"} {
		if _, err := ParseTimePeriod(s); err == nil {
			t.Errorf("accepted %q", s)
		}
	}
	if _, err := Parse("xx/01/2020:01/02/2020"); err == nil {
		t.Error("lost start parse error")
	}
}

func FuzzParseTimePeriod(f *testing.F) {
	for _, s := range []string{"", "garbage", "2020-01-01", "2020-01-01 -> 2020-01-02", "2020-01-01 2020-01-02 2020-01-03"} {
		f.Add(s)
	}
	f.Fuzz(func(t *testing.T, s string) {
		period, err := ParseTimePeriod(s)
		if err == nil && period.End.Before(period.Start) {
			t.Fatal("accepted reversed range")
		}
	})
}

func TestRegressionSmoothingOnePeriod(t *testing.T) {
	for _, factory := range []func(Indicator, int) Indicator{NewEMAIndicator, NewMMAIndicator} {
		ind := factory(NewFixedIndicator(math.Inf(1), 2, 3), 1)
		auditEqual(t, ind.Calculate(2), 3)
		if !math.IsInf(ind.Calculate(0).Float(), 1) {
			t.Fatal("lost infinity")
		}
		auditEqual(t, ind.Calculate(1), 2)
	}
}

func TestRegressionSmoothingIgnoresJSONFormatting(t *testing.T) {
	before := big.MarshalQuoted
	big.MarshalQuoted = true
	defer func() { big.MarshalQuoted = before }()
	s := mockTimeSeries("0.12345678901234567890123456789", "0.12345678901234567890123456789")
	got := NewEMAIndicator(NewClosePriceIndicator(s), 2).Calculate(1)
	if got.Sub(s.Candles[0].ClosePrice).Abs().GT(big.NewFromString("1e-70")) {
		t.Fatal("quoted JSON setting changed precision")
	}
}

func TestRegressionCacheResetAfterReindexing(t *testing.T) {
	s := auditSeries(10, 12, 9, 11, 10, 12, 14, 13, 15, 12)
	makeIndicator := func() Indicator {
		return NewMACDHistogramIndicator(NewMACDIndicator(NewClosePriceIndicator(s), 2, 3), 2)
	}
	ind := makeIndicator()
	ind.Calculate(s.LastIndex())
	s.Candles = s.Candles[2:]
	ResetCacheFrom(ind, -1)
	fresh := makeIndicator()
	for i := 0; i <= s.LastIndex(); i++ {
		auditEqual(t, ind.Calculate(i), fresh.Calculate(i).Float())
	}
	// Reset beyond the cache is safe and does not invalidate existing results.
	before := ind.Calculate(s.LastIndex())
	ResetCacheFrom(ind, 10000)
	auditEqual(t, ind.Calculate(s.LastIndex()), before.Float())
}

func TestRegressionAroonUsesLatestEqualExtreme(t *testing.T) {
	ind := NewAroonUpIndicator(NewFixedIndicator(1, 5, 5, 4, 3, 2), 3)
	auditEqual(t, ind.Calculate(2), 100)
	auditEqual(t, ind.Calculate(4), 100.0/3)
	auditEqual(t, ind.Calculate(5), 100.0/3)
}

func TestRegressionStrategyBoundary(t *testing.T) {
	rule := OverIndicatorRule{First: NewConstantIndicator(2), Second: NewConstantIndicator(1)}
	strategy := RuleStrategy{EntryRule: rule, ExitRule: rule, UnstablePeriod: 5}
	record := NewTradingRecord()
	if strategy.ShouldEnter(4, record) || !strategy.ShouldEnter(5, record) {
		t.Error("incorrect entry boundary")
	}
	record.Operate(Order{Side: BUY, Price: big.ONE, Amount: big.ONE})
	if strategy.ShouldExit(4, record) || !strategy.ShouldExit(5, record) {
		t.Error("incorrect exit boundary")
	}
}

func TestRegressionMixedPositionReturns(t *testing.T) {
	r := auditRecord(BUY, 100, 110, time.Second)
	r.Operate(Order{Side: SELL, Price: big.NewFromInt(110), Amount: big.ONE, ExecutionTime: time.Unix(2, 0)})
	r.Operate(Order{Side: BUY, Price: big.NewFromInt(99), Amount: big.ONE, ExecutionTime: time.Unix(3, 0)})
	if got := (PercentGainAnalysis{}).Analyze(r); math.Abs(got-.21) > 1e-12 {
		t.Errorf("mixed return=%v want=0.21", got)
	}
	if got := (TotalProfitAnalysis{}).Analyze(r); got != 21 {
		t.Errorf("mixed profit=%v want=21", got)
	}
}
