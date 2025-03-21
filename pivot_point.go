package techan

import "github.com/sdcoffey/big"

// PivotPoint represents the pivot point and its support and resistance levels.
type PivotPoint struct {
	Pivot, R1, R2, R3, R4, R5, S1, S2, S3, S4, S5 big.Decimal
}

// NewPivotPoint calculates the pivot point and its support and resistance levels.
func NewPivotPoint(high, low, close big.Decimal) PivotPoint {
	pivot := high.Add(low).Add(close).Div(big.NewDecimal(3))

	r1 := pivot.Mul(big.NewDecimal(2)).Sub(low)
	r2 := pivot.Add(high.Sub(low))
	r3 := high.Add(pivot.Sub(low).Mul(big.NewDecimal(2)))
	r4 := high.Add(pivot.Sub(low).Mul(big.NewDecimal(3)))
	r5 := high.Add(pivot.Sub(low).Mul(big.NewDecimal(4)))

	s1 := pivot.Mul(big.NewDecimal(2)).Sub(high)
	s2 := pivot.Sub(high.Sub(low))
	s3 := low.Sub(high.Sub(pivot).Mul(big.NewDecimal(2)))
	s4 := low.Sub(high.Sub(pivot).Mul(big.NewDecimal(3)))
	s5 := low.Sub(high.Sub(pivot).Mul(big.NewDecimal(4)))

	return PivotPoint{
		Pivot: pivot,
		R1:    r1,
		R2:    r2,
		R3:    r3,
		R4:    r4,
		R5:    r5,
		S1:    s1,
		S2:    s2,
		S3:    s3,
		S4:    s4,
		S5:    s5,
	}
}
