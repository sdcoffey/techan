package talib4g

import (
	"testing"
)

func TestCrossIndicator(t *testing.T) {
	fixed := NewFixedIndicator(12, 11, 10, 9, 11, 8, 7, 6)
	cross := NewCrossIndicator(fixed, NewConstantIndicator(10))

	decimalEquals(t, 0, cross.Calculate(0))
	decimalEquals(t, 0, cross.Calculate(1))
	decimalEquals(t, 1, cross.Calculate(2))
	decimalEquals(t, 0, cross.Calculate(3))
	decimalEquals(t, 0, cross.Calculate(4))
	decimalEquals(t, 1, cross.Calculate(5))
	decimalEquals(t, 0, cross.Calculate(6))
	decimalEquals(t, 0, cross.Calculate(7))
}
