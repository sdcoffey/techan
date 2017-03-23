package test

import (
	"github.com/sdcoffey/talib4g/indicators"
	"github.com/shopspring/decimal"
	"testing"

	. "github.com/stretchr/testify/assert"
)

func TestCrossIndicator(t *testing.T) {
	fixed := indicators.NewFixedIndicator(1, 2, 3, 4, 5, 4, 3, 6)
	cross := indicators.CrossIndicator{
		Upper: indicators.ConstantIndicator(decimal.New(4, 0)),
		Lower: fixed,
	}

	True(t, cross.Calculate(0).Equal(decimal.Zero))
	True(t, cross.Calculate(1).Equal(decimal.Zero))
	True(t, cross.Calculate(2).Equal(decimal.Zero))
	True(t, cross.Calculate(3).Equal(decimal.Zero))
	True(t, cross.Calculate(4).Equal(decimal.New(1, 0)))
	True(t, cross.Calculate(5).Equal(decimal.Zero))
	True(t, cross.Calculate(6).Equal(decimal.Zero))
	True(t, cross.Calculate(7).Equal(decimal.New(1, 0)))
}
