package talib4g

import (
	"math/big"
)

type BigInt struct {
	rep []uint64
	neg bool
}

const (
	fullMask64 = ^uint64(0)
)

func Zero() BigInt {
	return BigInt{rep: []uint64{0}}
}

func NewBigInt(size int) BigInt {
	return BigInt{
		rep: make([]uint64, size/64),
		neg: false,
	}
}

func NewBigInt64(value int64) BigInt {
	return BigInt{
		neg: value < 0,
		rep: []uint64{uint64(Abs(int(value)))},
	}
}

// Returns this number as an integer value. If the integer in question
// is larger than 64 bits in size, only the first 64 bits will be returned
func (x BigInt) Int() int64 {
	var i64 int64
	if x.rep[0]&1<<64 > 0 {
		i64 = int64(x.rep[0] >> 1)
	} else {
		i64 = int64(x.rep[0])
	}

	if x.neg {
		return -i64
	}
	return i64
}

// Cmp compares x and y and returns:
//
//   -1 if x <  y
//    0 if x == y
//   +1 if x >  y
//
func (x BigInt) Cmp(y BigInt) int {
	if x.neg == y.neg {
		if len(x.rep) > len(y.rep) {
			return 1
		} else if len(y.rep) > len(x.rep) {
			return -1
		} else {
			for i := len(x.rep) - 1; i >= 0; i-- {
				if x.rep[i] > y.rep[i] {
					return 1
				} else if x.rep[i] < y.rep[i] {
					return -1
				}
			}
		}
	} else if x.neg {
		return -1
	} else {
		return 1
	}

	return 0
}

func (x BigInt) And(y BigInt) (nbi BigInt) {
	xl, yl := len(x.rep), len(y.rep)
	min := Min(xl, yl)

	if min == xl {
		nbi = x
	} else {
		nbi = y
	}

	for i := 0; i < min; i++ {
		nbi.rep[i] = x.rep[i] & y.rep[i]
	}

	return nbi.trim()
}

func (x BigInt) Or(y BigInt) (nbi BigInt) {
	xl, yl := len(x.rep), len(y.rep)
	max := Max(xl, yl)

	if max == xl {
		nbi = x
	} else {
		nbi = y
	}

	for i := max - 1; i >= 0; i-- {
		if i < xl && i < yl {
			nbi.rep[i] = x.rep[i] | y.rep[i]
		} else if i < xl {
			nbi.rep[i] = x.rep[i]
		} else if i < yl {
			nbi.rep[i] = y.rep[i]
		}
	}

	return nbi
}

func (x BigInt) Xor(y BigInt) BigInt {
	for i := 0; i < Max(len(x.rep), len(y.rep)); i++ {
		if i < len(x.rep) {
			if i < len(y.rep) {
				x.rep[i] = x.rep[i] ^ y.rep[i]
			} else {
				x.rep[i] = x.rep[i]
			}
		} else {
			x.rep = append(x.rep, y.rep[i])
		}
	}

	return x.trim()
}

func (x BigInt) Zero() bool {
	for i := len(x.rep) - 1; i >= 0; i-- {
		if x.rep[i]&fullMask64 > 0 {
			return false
		}
	}

	return true
}

func (x BigInt) Size() int {
	return len(x.rep) * 64
}

func (x BigInt) Add(y BigInt) BigInt {
	const lastBit = uint64(1 << 63)

	if !x.neg && y.neg {
		return x.Sub(y)
	} else if x.neg && !y.neg {
		switch x.Abs().Cmp(y) {
		case 0:
			return Zero()
		case -1:
			return Zero().Add(y.Sub(x.Abs()))
		case 1:
			return x.Abs().Sub(y).Neg()
		}
	}

	min := Min(len(x.rep), len(y.rep))
	for i := 0; i < min; i++ {
		xAt, yAt := x.rep[i], y.rep[i]
		if (xAt&yAt)&lastBit > 0 {
			x = x.LShift(1)
			x.rep[i] = (xAt ^ lastBit) + (yAt ^ lastBit)
		} else {
			x.rep[i] += yAt
		}
	}

	if len(y.rep) > len(x.rep) {
		x.rep = append(x.rep, y.rep[len(y.rep)-len(x.rep):]...)
	}

	return x.trim()
}

func (x BigInt) Sub(y BigInt) BigInt {
	if y.neg {
		return x.Add(y.Abs())
	} else if !x.neg { // x > 0
		switch x.Cmp(y) { // x == y
		case 0:
			return Zero()
		case -1: // x < y
			return Zero().Add(y.Abs().Sub(x)).Neg()
		}
	} else {
		return x.Abs().Add(y).Neg()
	}

	for i := 0; i < Max(len(x.rep), len(y.rep)); i++ {
		if i < len(x.rep) {
			if i < len(y.rep) {
				if x.rep[i] < y.rep[i] {
					x = x.RShift(1)
					x.rep[i] = fullMask64 - y.rep[i] + 1
				} else {
					x.rep[i] = x.rep[i] - y.rep[i]
				}
			} else {
				x.rep[i] = x.rep[i]
			}
		} else {
			x.rep = append(x.rep, y.rep[i])
		}
	}

	return x.trim()
}

func (x BigInt) Mul(y BigInt) BigInt {
	if x.Zero() || y.Zero() {
		return Zero()
	}

	xn, yn := x.neg, y.neg

	if x.neg {
		x = x.Abs()
	}

	if y.neg {
		y = y.Abs()
	}

	if x.Cmp(y) < 0 {
		tmp := y
		y = x
		x = tmp
	}

	product := Zero()
	for !y.Zero() {
		if y.rep[0]&1 > 0 {
			product = product.Add(x)
		}
		x = x.LShift(1)
		y = y.RShift(1)
	}

	if xn && yn {
		return product.Abs()
	} else if xn || yn {
		return product.Neg()
	}

	return product
}

var one = NewBigInt64(1)

func (x BigInt) Div(y BigInt) BigInt {
	if y.Zero() {
		return Zero()
	}

	xn, yn := x.neg, y.neg

	if x.neg {
		x = x.Abs()
	}

	if y.neg {
		y = y.Abs()
	}

	var mask = NewBigInt64(1)
	quo := Zero()

	for y.Cmp(x) <= 0 {
		y = y.LShift(1)
		mask = mask.LShift(1)
	}

	for mask.Cmp(one) > 0 {
		y = y.RShift(1)
		mask = mask.RShift(1)
		if x.Cmp(y) >= 0 {
			x = x.Sub(y)
			quo = quo.Or(mask)
		}
	}

	if xn && yn {
		return quo.Abs()
	} else if xn || yn {
		return quo.Neg()
	}

	return quo
}

func biOrder(x, y BigInt) (BigInt, BigInt) {
	if x.Cmp(y) < 0 {
		return x, y
	}

	return y, x
}

func (x BigInt) RShift(bits uint64) BigInt {
	for bits > 0 {
		s := -Min(64, int(bits))
		x.rep = shift(x.rep, s)
		bits += uint64(s)
	}

	return x
}

func (x BigInt) LShift(bits uint64) BigInt {
	for bits > 0 {
		s := Min(64, int(bits))
		x.rep = shift(x.rep, s)
		bits -= uint64(s)
	}

	return x
}

func (x BigInt) Neg() BigInt {
	x.neg = !x.neg

	return x
}

func (x BigInt) Abs() BigInt {
	if x.neg {
		x = x.Neg()
	}

	return x
}

func (x BigInt) HighestBit() uint64 {
	li := len(x.rep) - 1
	return highestBit(x.rep[li]) + uint64(li*64)
}

func highestBit(x uint64) uint64 {
	for j := 63; j >= 0; j-- {
		if x&(1<<uint64(j)) > 0 {
			return uint64(j)
		}
	}

	return 0
}

// TODO re-implement this in assembly
func shift(vals []uint64, numbits int) []uint64 {
	var mask uint64

	if numbits > 0 {
		mask = fullMask64 >> uint64(64-numbits) << uint64(64-numbits)
		for i := len(vals) - 1; i >= 0; i-- {
			maskResult := vals[i] & mask
			if maskResult > 0 {
				if i == len(vals)-1 {
					vals = append(vals, maskResult>>uint64(64-numbits))
					vals[len(vals)-2] <<= uint64(numbits)
				} else {
					vals[i] |= maskResult
				}
			} else {
				vals[i] <<= uint64(numbits)
			}
		}
	} else {
		mask = fullMask64 >> uint64(64+numbits)
		absshift := uint64(-numbits)
		for i := 0; i < len(vals); i++ {
			if i == 0 {
				vals[i] >>= absshift
			} else {
				maskResult := vals[len(vals)-1] & mask
				if maskResult > 0 {
					if i == len(vals)-1 {
						vals[i] >>= absshift
						vals[i-1] |= maskResult << uint64(64-absshift)
						if vals[i]&mask == 0 {
							vals = vals[:i]
						}
					}
				} else {
					vals[i] >>= absshift
				}
			}
		}
	}

	return vals
}

func (x BigInt) trim() BigInt {
	if len(x.rep) > 0 {
		lastIndex := len(x.rep)
		for i := lastIndex - 1; i > 0; i-- {
			if x.rep[i]&fullMask64 == 0 {
				lastIndex = i
			} else {
				break
			}
		}

		x.rep = x.rep[:lastIndex]
	}

	return x
}

func (x BigInt) BinaryString() string {
	return x.AsBig().Text(2)
}

func (x BigInt) AsBig() (bigInt *big.Int) {
	bigInt = big.NewInt(0)
	for i := 0; i < len(x.rep); i++ {
		carry := x.rep[i] & 1
		asi64 := int64(x.rep[i] >> 1)
		or := big.NewInt(asi64)
		or = or.Lsh(or, 1)
		or = or.Or(or, big.NewInt(int64(carry)))

		or = or.Lsh(or, uint(64*i))
		bigInt = bigInt.Or(bigInt, or)
	}

	if x.neg {
		bigInt = bigInt.Neg(bigInt)
	}

	return bigInt
}

func (x BigInt) String() string {
	return x.AsBig().String()
}
