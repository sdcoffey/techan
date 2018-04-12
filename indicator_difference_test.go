package techan

import (
	"testing"
)

func TestDifferenceIndicator_Calculate(t *testing.T) {
	di := NewDifferenceIndicator(NewFixedIndicator(10, 9, 8), NewFixedIndicator(8, 9, 10))

	decimalEquals(t, 2, di.Calculate(0))
	decimalEquals(t, 0, di.Calculate(1))
	decimalEquals(t, -2, di.Calculate(2))
}
