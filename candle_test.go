package techan

import (
	"testing"
	"time"

	"fmt"
	"strings"

	"github.com/sdcoffey/big"
	"github.com/stretchr/testify/assert"
)

func TestCandle_AddTrade(t *testing.T) {
	now := time.Now()
	candle := NewCandle(TimePeriod{
		Start: now,
		End:   now.Add(time.Minute),
	})

	candle.AddTrade(big.NewDecimal(1), big.NewDecimal(2)) // Open
	candle.AddTrade(big.NewDecimal(1), big.NewDecimal(5)) // High
	candle.AddTrade(big.NewDecimal(1), big.NewDecimal(1)) // Low
	candle.AddTrade(big.NewDecimal(1), big.NewDecimal(3)) // No Diff
	candle.AddTrade(big.NewDecimal(1), big.NewDecimal(3)) // Close

	assert.EqualValues(t, 2, candle.OpenPrice.Float())
	assert.EqualValues(t, 5, candle.MaxPrice.Float())
	assert.EqualValues(t, 1, candle.MinPrice.Float())
	assert.EqualValues(t, 3, candle.ClosePrice.Float())
	assert.EqualValues(t, 5, candle.Volume.Float())
	assert.EqualValues(t, 5, candle.TradeCount)
}

func TestCandle_String(t *testing.T) {
	now := time.Now()
	candle := NewCandle(TimePeriod{
		Start: now,
		End:   now.Add(time.Minute),
	})

	candle.ClosePrice = big.NewFromString("1")
	candle.OpenPrice = big.NewFromString("2")
	candle.MaxPrice = big.NewFromString("3")
	candle.MinPrice = big.NewFromString("0")
	candle.Volume = big.NewFromString("10")

	expected := strings.TrimSpace(fmt.Sprintf(`
Time:	%s
Open:	2.00
Close:	1.00
High:	3.00
Low:	0.00
Volume:	10.00
`, candle.Period))

	assert.EqualValues(t, expected, candle.String())
}
