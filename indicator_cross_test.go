package talib4g

import (
	"testing"
)

func TestCrossIndicator(t *testing.T) {
	fixed := NewFixedIndicator(8, 9, 10, 12, 9, 11, 12, 13, 10, 9)
	cross := NewCrossIndicator(fixed, NewConstantIndicator(10))

	decimalEquals(t, 0, cross.Calculate(0))
	decimalEquals(t, 0, cross.Calculate(1))
	decimalEquals(t, 0, cross.Calculate(2))
	decimalEquals(t, 0, cross.Calculate(3))
	decimalEquals(t, 1, cross.Calculate(4))
	decimalEquals(t, 0, cross.Calculate(5))
	decimalEquals(t, 0, cross.Calculate(6))
	decimalEquals(t, 0, cross.Calculate(7))
	decimalEquals(t, 0, cross.Calculate(8))
	decimalEquals(t, 1, cross.Calculate(9))
}
