package talib4g

import (
	"testing"
	"time"

	"bytes"

	"bufio"

	"fmt"

	"github.com/sdcoffey/big"
	"github.com/stretchr/testify/assert"
)

const example = "EXM"

func TestTotalProfitAnalysis(t *testing.T) {
	t.Run("Simple", func(t *testing.T) {
		record := NewTradingRecord()
		tpa := TotalProfitAnalysis{}

		record.Enter(big.NewDecimal(1), big.NewDecimal(1), big.ZERO, example, time.Now())
		record.Exit(big.NewDecimal(2), big.NewDecimal(1), big.ZERO, example, time.Now()) // Gain 1

		assert.EqualValues(t, 1, tpa.Analyze(record))

		record.Enter(big.NewDecimal(1), big.NewDecimal(1), big.ZERO, example, time.Now())
		record.Exit(big.NewDecimal(.5), big.NewDecimal(1), big.ZERO, example, time.Now()) // Lose .5

		assert.EqualValues(t, .5, tpa.Analyze(record))
	})
}

func TestPercentGainAnalysis(t *testing.T) {
	t.Run("Zero", func(t *testing.T) {
		record := NewTradingRecord()

		pga := PercentGainAnalysis{}

		assert.EqualValues(t, 0, pga.Analyze(record))
	})

	t.Run("Simple gain", func(t *testing.T) {
		record := NewTradingRecord()

		pga := PercentGainAnalysis{}

		record.Enter(big.NewDecimal(1), big.NewDecimal(1), big.ZERO, example, time.Now())
		record.Exit(big.NewDecimal(2), big.NewDecimal(1), big.ZERO, example, time.Now())

		gain := pga.Analyze(record)
		assert.EqualValues(t, 1, gain)
	})

	t.Run("Simple loss", func(t *testing.T) {
		record := NewTradingRecord()

		pga := PercentGainAnalysis{}

		record.Enter(big.NewDecimal(2), big.NewDecimal(1), big.ZERO, example, time.Now())
		record.Exit(big.NewDecimal(1), big.NewDecimal(1), big.ZERO, example, time.Now())

		gain := pga.Analyze(record)
		assert.EqualValues(t, -.5, gain)
	})

	t.Run("Small loss and gain", func(t *testing.T) {
		record := NewTradingRecord()

		pga := PercentGainAnalysis{}

		record.Enter(big.NewDecimal(2), big.NewDecimal(1), big.ZERO, example, time.Now())
		record.Exit(big.NewDecimal(1), big.NewDecimal(1), big.ZERO, example, time.Now())

		record.Enter(big.NewDecimal(1), big.NewDecimal(1), big.ZERO, example, time.Now())
		record.Exit(big.NewDecimal(1.25), big.NewDecimal(1), big.ZERO, example, time.Now())

		gain := pga.Analyze(record)
		assert.EqualValues(t, -.375, gain)
	})
}

func TestNumTradesAnalysis(t *testing.T) {
	record := NewTradingRecord()

	var nta NumTradesAnalysis

	assert.EqualValues(t, 0, nta.Analyze(record))
}

func TestLogTradesAnalysis(t *testing.T) {
	buffer := bytes.NewBufferString("")

	logger := LogTradesAnalysis{
		Writer: buffer,
	}

	record := NewTradingRecord()

	now := time.Now().UTC()
	dates := []time.Time{
		now,
		now.AddDate(0, 0, 1),
		now.AddDate(0, 0, 2),
		now.AddDate(0, 0, 3),
	}

	record.Enter(big.NewDecimal(2), big.NewDecimal(1), big.ZERO, example, dates[0])
	record.Exit(big.NewDecimal(1), big.NewDecimal(1), big.ZERO, example, dates[1])

	record.Enter(big.NewDecimal(1), big.NewDecimal(1), big.ZERO, example, dates[2])
	record.Exit(big.NewDecimal(1.25), big.NewDecimal(1), big.ZERO, example, dates[3])

	val := logger.Analyze(record)
	assert.EqualValues(t, 0, val)

	scanner := bufio.NewScanner(buffer)

	var i int
	for scanner.Scan() {
		text := scanner.Text()

		var expected string
		switch i {
		case 0:
			expected = fmt.Sprintf("%s - enter with buy EXM (1 @ $2)", dates[0].Format(time.RFC822))
			break
		case 1:
			expected = fmt.Sprintf("%s - exit with sell EXM (1 @ $1)", dates[1].Format(time.RFC822))
			break
		case 2:
			expected = fmt.Sprintf("Profit: $-1")
			break
		case 3:
			expected = fmt.Sprintf("%s - enter with buy EXM (1 @ $1)", dates[2].Format(time.RFC822))
			break
		case 4:
			expected = fmt.Sprintf("%s - exit with sell EXM (1 @ $1.25)", dates[3].Format(time.RFC822))
			break
		case 5:
			expected = "Profit: $0.25"
			break
		}

		assert.EqualValues(t, expected, text)
		i++
	}
}

func TestPeriodProfitAnalysis(t *testing.T) {
	record := NewTradingRecord()

	now := time.Now().Add(-time.Minute * 5)

	record.Enter(big.NewDecimal(1), big.NewDecimal(1), big.ZERO, example, now)
	record.Exit(big.NewDecimal(2), big.NewDecimal(1), big.ZERO, example, now.Add(time.Minute))

	record.Enter(big.NewDecimal(2), big.NewDecimal(1), big.ZERO, example, now.Add(time.Minute*2))
	record.Exit(big.NewDecimal(3), big.NewDecimal(1), big.ZERO, example, now.Add(time.Minute*3))

	ppa := PeriodProfitAnalysis{
		Period: time.Minute * 2,
	}

	assert.EqualValues(t, 2, ppa.Analyze(record))
}

func TestProfitableTradesAnalysis(t *testing.T) {
	record := NewTradingRecord()

	record.Enter(big.NewDecimal(1), big.NewDecimal(1), big.ZERO, example, time.Now())
	record.Exit(big.NewDecimal(2), big.NewDecimal(1), big.ZERO, example, time.Now())

	record.Enter(big.NewDecimal(2), big.NewDecimal(1), big.ZERO, example, time.Now())
	record.Exit(big.NewDecimal(1), big.NewDecimal(1), big.ZERO, example, time.Now())

	pta := ProfitableTradesAnalysis{}

	assert.EqualValues(t, 1, pta.Analyze(record))
}

func TestAverageProfitAnalysis(t *testing.T) {
	record := NewTradingRecord()

	record.Enter(big.NewDecimal(1), big.NewDecimal(1), big.ZERO, example, time.Now())
	record.Exit(big.NewDecimal(2), big.NewDecimal(1), big.ZERO, example, time.Now())

	record.Enter(big.NewDecimal(2), big.NewDecimal(1), big.ZERO, example, time.Now())
	record.Exit(big.NewDecimal(5), big.NewDecimal(1), big.ZERO, example, time.Now())

	pta := AverageProfitAnalysis{}

	assert.EqualValues(t, 2, pta.Analyze(record))
}

func TestBuyAndHoldAnalysis(t *testing.T) {
	series := mockTimeSeries("1", "2", "3", "2", "6")
	record := NewTradingRecord()

	t.Run("== 0 trades returns zero", func(t *testing.T) {
		buyAndHoldAnalysis := BuyAndHoldAnalysis{
			TimeSeries:    series,
			StartingMoney: 1,
		}

		assert.EqualValues(t, 0, buyAndHoldAnalysis.Analyze(record))
	})

	t.Run("> 0 trades", func(t *testing.T) {
		record.Enter(big.NewDecimal(1), big.NewDecimal(1), big.ZERO, example, time.Now())
		record.Exit(big.NewDecimal(2), big.NewDecimal(1), big.ZERO, example, time.Now())

		record.Enter(big.NewDecimal(3), big.NewDecimal(1), big.ZERO, example, time.Now())
		record.Exit(big.NewDecimal(6), big.NewDecimal(1), big.ZERO, example, time.Now())

		buyAndHoldAnalysis := BuyAndHoldAnalysis{
			TimeSeries:    series,
			StartingMoney: 1,
		}

		assert.EqualValues(t, 5, buyAndHoldAnalysis.Analyze(record))
	})
}
