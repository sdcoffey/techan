package techan

import (
	"math"
	"testing"
)

func TestFixedIndicator_Calculate(t *testing.T) {
	fi := NewFixedIndicator(0, 1, 2, -100, math.MaxInt64)

	decimalEquals(t, 0, fi.Calculate(0))
	decimalEquals(t, 1, fi.Calculate(1))
	decimalEquals(t, 2, fi.Calculate(2))
	decimalEquals(t, -100, fi.Calculate(3))
	decimalEquals(t, math.MaxInt64, fi.Calculate(4))
}
