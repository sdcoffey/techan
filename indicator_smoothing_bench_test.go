package techan

import (
	"fmt"
	"testing"

	"github.com/sdcoffey/big"
)

var smoothingBenchmarkResult big.Decimal

func BenchmarkSmoothingHistory(b *testing.B) {
	for _, n := range []int{400, 800, 1600} {
		values := make([]float64, n)
		for i := range values {
			values[i] = float64(100 + i%7)
		}
		base := NewFixedIndicator(values...)
		for name, factory := range map[string]func(Indicator, int) Indicator{"EMA": NewEMAIndicator, "MMA": NewMMAIndicator} {
			b.Run(fmt.Sprintf("%s/%d", name, n), func(b *testing.B) {
				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					smoothingBenchmarkResult = factory(base, 10).Calculate(n - 1)
				}
			})
		}
	}
}
