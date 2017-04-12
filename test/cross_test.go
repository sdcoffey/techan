package test

import (
	. "github.com/sdcoffey/talib4g"
	"testing"

	. "github.com/stretchr/testify/assert"
)

func TestCrossIndicator(t *testing.T) {
	fixed := NewFixedIndicator(12, 11, 10, 9, 11, 8, 7, 6)
	cross := CrossIndicator{
		Lower: ConstantIndicator(10.0),
		Upper: fixed,
	}

	EqualValues(t, 0.0, cross.Calculate(0))
	EqualValues(t, 0.0, cross.Calculate(1))
	EqualValues(t, 0.0, cross.Calculate(2))
	EqualValues(t, 1.0, cross.Calculate(3))
	EqualValues(t, 0.0, cross.Calculate(4))
	EqualValues(t, 1.0, cross.Calculate(5))
	EqualValues(t, 0.0, cross.Calculate(6))
	EqualValues(t, 0.0, cross.Calculate(7))
}
