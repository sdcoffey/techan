package talib4g

import (
	. "github.com/stretchr/testify/assert"
	"testing"
)

func TestCrossIndicator(t *testing.T) {
	fixed := NewFixedIndicator(12, 11, 10, 9, 11, 8, 7, 6)
	cross := NewCrossIndicator(fixed, NewConstantIndicator(10))

	EqualValues(t, 0, cross.Calculate(0))
	EqualValues(t, 0, cross.Calculate(1))
	EqualValues(t, 1, cross.Calculate(2))
	EqualValues(t, 0, cross.Calculate(3))
	EqualValues(t, 0, cross.Calculate(4))
	EqualValues(t, 1, cross.Calculate(5))
	EqualValues(t, 0, cross.Calculate(6))
	EqualValues(t, 0, cross.Calculate(7))
}
