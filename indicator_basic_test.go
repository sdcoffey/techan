package talib4g

import (
	"testing"
	"time"

	"github.com/sdcoffey/big"
	"github.com/stretchr/testify/assert"
)

func TestTypicalPriceIndicator_Calculate(t *testing.T) {
	series := NewTimeSeries()

	candle := NewCandle(TimePeriod{
		Start: time.Now(),
		End:   time.Now().Add(time.Minute),
	})
	candle.MinPrice = big.NewFromString("1.2080")
	candle.MaxPrice = big.NewFromString("1.22")
	candle.ClosePrice = big.NewFromString("1.215")

	series.AddCandle(candle)

	typicalPrice := NewTypicalPriceIndicator(series).Calculate(0)

	assert.EqualValues(t, "1.2143", typicalPrice.FormattedString(4))
}
