package techan

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPow(t *testing.T) {
	assert.EqualValues(t, 1024, Pow(4, 5))
}

func TestAbs(t *testing.T) {
	t.Run("Positive", func(t *testing.T) {
		assert.EqualValues(t, 100, Abs(100))
	})

	t.Run("Negative", func(t *testing.T) {
		assert.EqualValues(t, 100, Abs(-100))
	})
}

func BenchmarkPower(b *testing.B) {
	var x int
	for i := 0; i < b.N; i++ {
		x = Pow(4, 5)
	}
	x++
}

func BenchmarkAbs(b *testing.B) {
	var x int
	for i := 0; i < b.N; i++ {
		x = Abs(-5)
	}
	x++
}
