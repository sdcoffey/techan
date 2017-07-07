package talib4g

import (
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

func TestNewBigInt64(t *testing.T) {
	t.Run("Positive", func(t *testing.T) {
		bi := NewBigInt64(8)

		assert.EqualValues(t, "8", bi.String())
	})

	t.Run("Negative", func(t *testing.T) {
		bi := NewBigInt64(-8)
		assert.EqualValues(t, "-8", bi.String())
	})
}

func TestBigInt_Int(t *testing.T) {
	bi := NewBigInt64(371282373)

	assert.EqualValues(t, "371282373", bi.String())
}

func TestBigInt_Cmp(t *testing.T) {
	t.Run("Same bit single depth eq", func(t *testing.T) {
		x := NewBigInt64(1)
		y := NewBigInt64(1)

		assert.EqualValues(t, 0, x.Cmp(y))
	})

	t.Run("Same bit depth lt", func(t *testing.T) {
		x := NewBigInt64(-1)
		y := NewBigInt64(1)

		assert.EqualValues(t, -1, x.Cmp(y))
	})

	t.Run("Same bit depth gt", func(t *testing.T) {
		x := NewBigInt64(1)
		y := NewBigInt64(-1)

		assert.EqualValues(t, 1, x.Cmp(y))
	})

	t.Run("Same bit multiple depth eq", func(t *testing.T) {
		x := NewBigInt64(1).LShift(64)
		y := NewBigInt64(1).LShift(64)

		assert.EqualValues(t, 0, x.Cmp(y))
	})

	t.Run("Different bit depth lhs", func(t *testing.T) {
		x := NewBigInt64(1).LShift(64)
		y := NewBigInt64(1).LShift(4)

		assert.EqualValues(t, 1, x.Cmp(y))
	})

	t.Run("Different bit depth rhs", func(t *testing.T) {
		x := NewBigInt64(1).LShift(4)
		y := NewBigInt64(1).LShift(54)

		assert.EqualValues(t, -1, x.Cmp(y))
	})

	t.Run("Neg cmp pos", func(t *testing.T) {
		x := NewBigInt64(-1)
		y := NewBigInt64(1)

		assert.EqualValues(t, -1, x.Cmp(y))
	})

	t.Run("Pos cmp neg", func(t *testing.T) {
		x := NewBigInt64(1)
		y := NewBigInt64(-1)

		assert.EqualValues(t, 1, x.Cmp(y))
	})
}

func TestBigInt_And(t *testing.T) {
	t.Run("Same bit depth", func(t *testing.T) {
		bi := NewBigInt64(5)
		bi = bi.And(NewBigInt64(1))

		assert.EqualValues(t, "1", bi.String())
	})

	t.Run("Different bit depths rgs", func(t *testing.T) {
		bi := NewBigInt64(1).LShift(64).Add(NewBigInt64(1024))

		assert.EqualValues(t, "1024", bi.And(NewBigInt64(1024)).String())
	})

	t.Run("Different bit depths lhs", func(t *testing.T) {
		x := NewBigInt64(1024)
		y := NewBigInt64(1024).LShift(64).Add(NewBigInt64(1024))

		assert.EqualValues(t, "1024", x.And(y).String())
	})
}

func TestBigInt_Or(t *testing.T) {
	t.Run("Same bit depth", func(t *testing.T) {
		bi := NewBigInt64(10)
		bi = bi.Or(NewBigInt64(1))

		assert.EqualValues(t, "11", bi.String())
	})

	t.Run("Different bit depths rhs", func(t *testing.T) {
		bi := NewBigInt64(1).LShift(64)

		assert.EqualValues(t, "18446744073709551617", bi.Or(NewBigInt64(1)).String())
	})

	t.Run("Different bit depths lhs", func(t *testing.T) {
		x := NewBigInt64(1024)
		y := NewBigInt64(1024).LShift(64).Add(NewBigInt64(1024))

		assert.EqualValues(t, "18889465931478580855808", x.Or(y).String())
	})
}

func TestBigInt_LShift(t *testing.T) {
	t.Run("Under 64 bits", func(t *testing.T) {
		bi := NewBigInt64(1).LShift(1)

		assert.EqualValues(t, "2", bi.String())
	})

	t.Run("Expand", func(t *testing.T) {
		bi := NewBigInt64(1 << 60)
		bi = bi.LShift(4)

		assert.EqualValues(t, "10000000000000000000000000000000000000000000000000000000000000000", bi.BinaryString())
	})

	t.Run("Expand size 2", func(t *testing.T) {
		bi := NewBigInt64(1 << 60)
		bi = bi.LShift(4)
		bi = bi.LShift(64)

		assert.EqualValues(t, "100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000", bi.BinaryString())
	})
}

func TestBigInt_RShift(t *testing.T) {
	t.Run("Under 64 bits", func(t *testing.T) {
		bi := NewBigInt64(1024)
		bi = bi.RShift(10)

		assert.EqualValues(t, "1", bi.BinaryString())
	})

	t.Run("Over 64 bits carry", func(t *testing.T) {
		bi := NewBigInt64(1 << 62)
		bi = bi.LShift(4)
		assert.Len(t, bi.rep, 2)

		bi = bi.RShift(4)

		assert.EqualValues(t, 1, len(bi.rep))
		assert.EqualValues(t, "100000000000000000000000000000000000000000000000000000000000000", bi.BinaryString())
	})

	t.Run("Over 64 bits no carry", func(t *testing.T) {
		bi := NewBigInt64(1 << 62).Or(NewBigInt64(1))
		bi = bi.LShift(10)
		bi = bi.RShift(4)

		assert.EqualValues(t, 2, len(bi.rep))
		assert.EqualValues(t, "100000000000000000000000000000000000000000000000000000000000001000000", bi.BinaryString())
	})
}

func TestBigInt_Xor(t *testing.T) {
	t.Run("Same bit depth", func(t *testing.T) {
		x := NewBigInt64(5)
		y := NewBigInt64(3)
		x = x.Xor(y)

		assert.EqualValues(t, "6", x.String())
	})

	t.Run("Different bit depths rhs", func(t *testing.T) {
		x := NewBigInt64(1).LShift(64)
		assert.EqualValues(t, "18446744073709551617", x.Xor(NewBigInt64(1)).String())
	})

	t.Run("Different bit depths lhs", func(t *testing.T) {
		x := NewBigInt64(1024)
		y := NewBigInt64(1024).LShift(64).Add(NewBigInt64(1024))

		assert.EqualValues(t, "18889465931478580854784", x.Xor(y).String())
	})
}

func TestBigInt_Add(t *testing.T) {
	t.Run("Same bit length", func(t *testing.T) {
		bi := NewBigInt64(10)
		bi = bi.Add(NewBigInt64(22))

		assert.EqualValues(t, "32", bi.String())
	})

	t.Run("Carry", func(t *testing.T) {
		x := NewBigInt64(1).LShift(63)
		y := NewBigInt64(1).LShift(63)

		res := x.Add(y)
		assert.EqualValues(t, 65, len(res.BinaryString()))
		assert.EqualValues(t, "18446744073709551616", res.String())
	})

	t.Run("More than one word", func(t *testing.T) {
		bi := NewBigInt64(1).LShift(64)
		bi = bi.Add(NewBigInt64(1).LShift(64))

		assert.EqualValues(t, "36893488147419103232", bi.String())
	})

	t.Run("More than one word rhs", func(t *testing.T) {
		bi := NewBigInt64(1).LShift(64)
		bi = bi.Add(NewBigInt64(10))

		assert.EqualValues(t, "18446744073709551626", bi.String())
	})

	t.Run("More than one word rhs", func(t *testing.T) {
		bi := NewBigInt64(10)
		bi = bi.Add(NewBigInt64(1).LShift(64))

		assert.EqualValues(t, "18446744073709551626", bi.String())
	})

	t.Run("negative add positive gt0", func(t *testing.T) {
		bi := NewBigInt64(-1)
		bi = bi.Add(NewBigInt64(2))

		assert.EqualValues(t, "1", bi.String())
	})

	t.Run("negative add positive lt0", func(t *testing.T) {
		bi := NewBigInt64(-10)
		bi = bi.Add(NewBigInt64(2))

		assert.EqualValues(t, "-8", bi.String())
	})

	t.Run("negative add negative lt0", func(t *testing.T) {
		bi := NewBigInt64(-10)
		bi = bi.Add(NewBigInt64(-2))

		assert.EqualValues(t, "-12", bi.String())
	})
}

func TestBigInt_Sub(t *testing.T) {
	t.Run("Same bit length", func(t *testing.T) {
		bi := NewBigInt64(10)
		bi = bi.Sub(NewBigInt64(2))

		assert.EqualValues(t, "8", bi.String())
	})

	t.Run("More than one word", func(t *testing.T) {
		bi := NewBigInt64(1).LShift(67)
		bi = bi.Sub(NewBigInt64(1).LShift(64))

		assert.EqualValues(t, "129127208515966861312", bi.String())
	})

	t.Run("More than one word lhs", func(t *testing.T) {
		bi := NewBigInt64(1).LShift(64)
		bi = bi.Sub(NewBigInt64(10))

		assert.EqualValues(t, "18446744073709551606", bi.String())
	})

	t.Run("More than one word rhs", func(t *testing.T) {
		bi := NewBigInt64(10)
		bi = bi.Sub(NewBigInt64(1).LShift(64))

		assert.EqualValues(t, "-18446744073709551606", bi.String())
	})

	t.Run("negative sub positive", func(t *testing.T) {
		bi := NewBigInt64(-1)
		bi = bi.Sub(NewBigInt64(2))

		assert.EqualValues(t, "-3", bi.String())
	})

	t.Run("negative sub negative", func(t *testing.T) {
		bi := NewBigInt64(-10)
		bi = bi.Sub(NewBigInt64(-11))

		assert.EqualValues(t, "1", bi.String())
	})

	t.Run("positive sub negative", func(t *testing.T) {
		bi := NewBigInt64(10)
		bi = bi.Sub(NewBigInt64(-2))

		assert.EqualValues(t, "12", bi.String())
	})
}

func TestBigInt_Mul(t *testing.T) {
	t.Run("Same bit length", func(t *testing.T) {
		x := NewBigInt64(4)
		y := NewBigInt64(8)

		assert.EqualValues(t, "32", x.Mul(y).String())
	})

	t.Run("More than one word lhs", func(t *testing.T) {
		x := NewBigInt64(1).LShift(64)
		y := NewBigInt64(372832)

		assert.EqualValues(t, "6877536486489279548096512", x.Mul(y).String())
	})

	t.Run("More than one word rhs", func(t *testing.T) {
		x := NewBigInt64(1).LShift(64)
		y := NewBigInt64(372832)

		assert.EqualValues(t, "6877536486489279548096512", y.Mul(x).String())
	})

	t.Run("neg mul pos", func(t *testing.T) {
		x := NewBigInt64(-8)
		y := NewBigInt64(8)

		assert.EqualValues(t, "-64", y.Mul(x).String())
	})

	t.Run("neg mul neg", func(t *testing.T) {
		x := NewBigInt64(-8)
		y := NewBigInt64(-8)

		assert.EqualValues(t, "64", y.Mul(x).String())
	})

	t.Run("one operand zero", func(t *testing.T) {
		x := NewBigInt64(-8)
		y := NewBigInt64(0)

		assert.EqualValues(t, "0", y.Mul(x).String())
	})
}

func TestBigInt_Div(t *testing.T) {
	t.Run("Same bit length no remainder", func(t *testing.T) {
		x := NewBigInt64(8)
		y := NewBigInt64(4)

		assert.EqualValues(t, "2", x.Div(y).String())
	})

	t.Run("Same bit length remainder", func(t *testing.T) {
		x := NewBigInt64(10)
		y := NewBigInt64(3)

		assert.EqualValues(t, "3", x.Div(y).String())
	})

	t.Run("Same bit length remainder", func(t *testing.T) {
		x := NewBigInt64(10)
		y := NewBigInt64(3)

		assert.EqualValues(t, "3", x.Div(y).String())
	})

	t.Run("Divide by self", func(t *testing.T) {
		x := NewBigInt64(10)
		y := NewBigInt64(10)

		assert.EqualValues(t, "1", x.Div(y).String())
	})

	t.Run("Divide by zero", func(t *testing.T) {
		x := NewBigInt64(10)
		y := NewBigInt64(0)

		assert.EqualValues(t, "0", x.Div(y).String())
	})

	t.Run("Neg div pos", func(t *testing.T) {
		x := NewBigInt64(-10)
		y := NewBigInt64(5)

		assert.EqualValues(t, "-2", x.Div(y).String())
	})

	t.Run("Pos div neg", func(t *testing.T) {
		x := NewBigInt64(10)
		y := NewBigInt64(-5)

		assert.EqualValues(t, "-2", x.Div(y).String())
	})

	t.Run("Neg div neg", func(t *testing.T) {
		x := NewBigInt64(-10)
		y := NewBigInt64(-5)

		assert.EqualValues(t, "2", x.Div(y).String())
	})
}

func TestBigInt_AsBig(t *testing.T) {
	t.Run("One word positive", func(t *testing.T) {
		b := NewBigInt64(100)
		asBigint := b.AsBig()
		assert.EqualValues(t, "100", asBigint.String())
	})

	t.Run("One word negative", func(t *testing.T) {
		b := NewBigInt64(-100)
		asBigint := b.AsBig()
		assert.EqualValues(t, "-100", asBigint.String())
	})

	t.Run("More than one word positive", func(t *testing.T) {
		b := NewBigInt64(1).LShift(100)
		asBigint := b.AsBig()
		assert.EqualValues(t, "1267650600228229401496703205376", asBigint.String())
	})

	t.Run("More than one word negative", func(t *testing.T) {
		b := NewBigInt64(-1).LShift(100)
		asBigint := b.AsBig()
		assert.EqualValues(t, "-1267650600228229401496703205376", asBigint.String())
	})

	t.Run("64th bit", func(t *testing.T) {
		b := NewBigInt64(1).LShift(63)
		asBigint := b.AsBig()
		assert.EqualValues(t, "9223372036854775808", asBigint.String())
	})
}

func TestBigInt_HighestBit(t *testing.T) {
	t.Run("One word", func(t *testing.T) {
		x := NewBigInt64(8)
		assert.EqualValues(t, 3, x.HighestBit())
	})

	t.Run("Two word", func(t *testing.T) {
		x := NewBigInt64(1).LShift(64)
		assert.EqualValues(t, 64, x.HighestBit())
	})
}

func TestBigInt_trim(t *testing.T) {
	b := NewBigInt64(0)
	b.rep = append(b.rep, 0)

	b = b.trim()
	assert.Len(t, b.rep, 1)
	assert.EqualValues(t, 0, b.rep[0])
}

// Benchmarks

func BenchmarkNewBigInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewBigInt64(371282373)
	}
}

func BenchmarkBigInt_Xor(b *testing.B) {
	bi := NewBigInt64(5)
	x := NewBigInt64(3)

	for i := 0; i < b.N; i++ {
		bi = bi.Xor(x)
	}
}

// allows discarding references in benchmarks without
func noop(bi BigInt) {}

func BenchmarkBigInt_Add(b *testing.B) {
	b.Run("Basic", func(b *testing.B) {
		bi := NewBigInt64(5)
		x := NewBigInt64(3)

		z := NewBigInt64(0)
		for i := 0; i < b.N; i++ {
			z = bi.Add(x)
		}
		noop(z)
	})

	b.Run("Carry", func(b *testing.B) {
		x := NewBigInt64(1).LShift(63)
		y := NewBigInt64(1).LShift(63)

		z := NewBigInt64(0)
		for i := 0; i < b.N; i++ {
			x.Add(y)
		}
		noop(z)
	})
}

func BenchmarkBigInt_Mul(b *testing.B) {
	b.Run("Basic", func(b *testing.B) {
		x := NewBigInt64(5)
		y := NewBigInt64(3)

		z := NewBigInt64(0)
		for i := 0; i < b.N; i++ {
			z = x.Mul(y)
		}
		noop(z) // Allows us to effectively discard this reference
	})

	b.Run("Multi word", func(b *testing.B) {
		x := NewBigInt64(5).LShift(64)
		y := NewBigInt64(3)

		z := NewBigInt64(0)
		for i := 0; i < b.N; i++ {
			z = x.Mul(y)
		}
		noop(z) // Allows us to effectively discard this reference
	})
}

func BenchmarkBigDotInt_Mul(b *testing.B) {
	b.Run("Basic", func(b *testing.B) {
		x := big.NewInt(10)
		y := big.NewInt(30)

		z := big.NewInt(1)
		for i := 0; i < b.N; i++ {
			z.Mul(x, y)
		}
	})

	b.Run("Multi word", func(b *testing.B) {
		x := big.NewInt(10)
		x = x.Lsh(x, 64)
		y := big.NewInt(30)

		z := big.NewInt(1)
		for i := 0; i < b.N; i++ {
			z.Mul(x, y)
		}
	})
}

func BenchmarkBigInt_Div(b *testing.B) {
	b.Run("Basic", func(b *testing.B) {
		x := NewBigInt64(5)
		y := NewBigInt64(3)

		z := NewBigInt64(0)
		for i := 0; i < b.N; i++ {
			z = x.Div(y)
		}
		noop(z) // Allows us to effectively discard this reference
	})

	b.Run("Multi word", func(b *testing.B) {
		x := NewBigInt64(5).LShift(64)
		y := NewBigInt64(3)

		z := NewBigInt64(0)
		for i := 0; i < b.N; i++ {
			z = x.Div(y)
		}
		noop(z) // Allows us to effectively discard this reference
	})
}

func BenchmarkBigDotInt_Div(b *testing.B) {
	b.Run("Basic", func(b *testing.B) {
		x := big.NewInt(10)
		y := big.NewInt(30)

		z := big.NewInt(1)
		for i := 0; i < b.N; i++ {
			z.Div(x, y)
		}
	})

	b.Run("Multi word", func(b *testing.B) {
		x := big.NewInt(10)
		x = x.Lsh(x, 64)
		y := big.NewInt(30)

		z := big.NewInt(1)
		for i := 0; i < b.N; i++ {
			z.Div(x, y)
		}
	})
}

func BenchmarkBigInt_Cmp(b *testing.B) {
	b.Run("Basic", func(b *testing.B) {
		x := NewBigInt64(10)
		y := NewBigInt64(8)

		for i := 0; i < b.N; i++ {
			x.Cmp(y)
		}
	})

	b.Run("Multi word", func(b *testing.B) {
		x := NewBigInt64(10).LShift(64)
		y := NewBigInt64(8).LShift(64)

		for i := 0; i < b.N; i++ {
			x.Cmp(y)
		}
	})
}

func BenchmarkBigDotInt_Cmp(b *testing.B) {
	b.Run("Basic", func(b *testing.B) {
		x := big.NewInt(10)
		y := big.NewInt(30)

		for i := 0; i < b.N; i++ {
			x.Cmp(y)
		}
	})

	b.Run("Multi word", func(b *testing.B) {
		x := big.NewInt(10)
		x = x.Lsh(x, 64)
		y := big.NewInt(30)
		y = y.Lsh(y, 64)

		for i := 0; i < b.N; i++ {
			x.Cmp(y)
		}
	})
}

func BenchmarkBigDotInt_Add(b *testing.B) {
	b.Run("Basic", func(b *testing.B) {
		x := big.NewInt(10)
		y := big.NewInt(30)

		z := big.NewInt(1)
		for i := 0; i < b.N; i++ {
			z.Add(x, y)
		}
	})

	b.Run("Carry", func(b *testing.B) {
		x := big.NewInt(1)
		x = x.Lsh(x, 63)
		y := big.NewInt(1)
		y = y.Lsh(y, 63)

		z := big.NewInt(1)
		for i := 0; i < b.N; i++ {
			z = z.Add(y, x)
		}
	})
}

func BenchmarkBigInt_Or(b *testing.B) {
	b.Run("Basic", func(b *testing.B) {
		bi := NewBigInt64(5)
		x := NewBigInt64(3)

		z := NewBigInt64(0)
		for i := 0; i < b.N; i++ {
			z = bi.Or(x)
		}
		noop(z) // Allows us to effectively discard this reference
	})

	b.Run("Multi word", func(b *testing.B) {
		bi := NewBigInt64(5)
		x := NewBigInt64(3).LShift(64)

		z := NewBigInt64(0)
		for i := 0; i < b.N; i++ {
			z = bi.Or(x)
		}
		noop(z) // Allows us to effectively discard this reference
	})
}

func BenchmarkBigDotInt_Or(b *testing.B) {
	b.Run("Basic", func(b *testing.B) {
		x := big.NewInt(5)
		y := big.NewInt(3)

		for i := 0; i < b.N; i++ {
			x.Or(x, y)
		}
	})

	b.Run("Multi word", func(b *testing.B) {
		x := big.NewInt(5)
		y := big.NewInt(3)
		y = y.Lsh(y, 64)

		for i := 0; i < b.N; i++ {
			x.Or(x, y)
		}
	})
}

func BenchmarkBigInt_LShift(b *testing.B) {
	b.Run("Basic", func(b1 *testing.B) {
		bi := NewBigInt64(1)

		b1.ResetTimer()
		for i := 0; i < b1.N; i++ {
			bi.LShift(10)
		}
	})

	b.Run("Multi word", func(b *testing.B) {
		bi := NewBigInt64(1)
		for i := 0; i < b.N; i++ {
			bi.LShift(64)
		}
	})
}

func BenchmarkBigDotInt_Lsh(b *testing.B) {
	b.Run("Basic", func(b *testing.B) {
		x := big.NewInt(1)

		z := big.NewInt(0)
		for i := 0; i < b.N; i++ {
			z = z.Lsh(x, 10)
		}
	})

	b.Run("Multi word", func(b *testing.B) {
		x := big.NewInt(1)

		z := big.NewInt(0)
		for i := 0; i < b.N; i++ {
			z = z.Lsh(x, 64)
		}
	})
}
