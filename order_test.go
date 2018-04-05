package talib4g

import (
	"testing"

	"github.com/sdcoffey/big"
	"github.com/stretchr/testify/assert"
)

func TestOrder_TotalAmount(t *testing.T) {
	t.Run("with commission", func(t *testing.T) {
		t.Run("buy", func(t *testing.T) {
			order := Order{
				Side:          BUY,
				Amount:        big.NewFromString("1"),
				Price:         big.NewFromString("10"),
				FeePercentage: big.NewFromString("0.01"),
			}

			assert.EqualValues(t, "0.99", order.TotalAmount().FormattedString(2))
		})

		t.Run("sell", func(t *testing.T) {
			order := Order{
				Side:          SELL,
				Amount:        big.NewFromString("1"),
				Price:         big.NewFromString("10"),
				FeePercentage: big.NewFromString("0.01"),
			}

			assert.EqualValues(t, "1.00", order.TotalAmount().FormattedString(2))
		})
	})

	t.Run("without commission", func(t *testing.T) {
		t.Run("buy", func(t *testing.T) {
			order := Order{
				Side:          BUY,
				Amount:        big.NewFromString("1"),
				Price:         big.NewFromString("10"),
				FeePercentage: big.NewFromString("0"),
			}

			assert.EqualValues(t, "1.00", order.TotalAmount().FormattedString(2))
		})

		t.Run("sell", func(t *testing.T) {
			order := Order{
				Side:   SELL,
				Amount: big.NewFromString("1"),
				Price:  big.NewFromString("10"),
			}

			assert.EqualValues(t, "1.00", order.TotalAmount().FormattedString(2))
		})
	})
}

func TestOrder_Cost(t *testing.T) {
	t.Run("with commission", func(t *testing.T) {
		t.Run("buy", func(t *testing.T) {
			order := Order{
				Side:          BUY,
				Amount:        big.NewFromString("1"),
				Price:         big.NewFromString("10"),
				FeePercentage: big.NewFromString("0.01"),
			}

			assert.EqualValues(t, "10.00", order.Cost().FormattedString(2))
		})

		t.Run("sell", func(t *testing.T) {
			order := Order{
				Side:          SELL,
				Amount:        big.NewFromString("1"),
				Price:         big.NewFromString("10"),
				FeePercentage: big.NewFromString("0.01"),
			}

			assert.EqualValues(t, "-10.00", order.Cost().FormattedString(2))
		})
	})

	t.Run("without commission", func(t *testing.T) {
		t.Run("buy", func(t *testing.T) {
			order := Order{
				Side:   BUY,
				Amount: big.NewFromString("1"),
				Price:  big.NewFromString("10"),
			}

			assert.EqualValues(t, "10.00", order.Cost().FormattedString(2))
		})

		t.Run("sell", func(t *testing.T) {
			order := Order{
				Side:          SELL,
				Amount:        big.NewFromString("1"),
				Price:         big.NewFromString("10"),
				FeePercentage: big.NewFromString("0.01"),
			}

			assert.EqualValues(t, "-10.00", order.Cost().FormattedString(2))
		})
	})
}

func TestOrder_Profit(t *testing.T) {
	t.Run("with commission", func(t *testing.T) {
		t.Run("buy", func(t *testing.T) {
			order := Order{
				Side:          BUY,
				Amount:        big.NewFromString("1"),
				Price:         big.NewFromString("10"),
				FeePercentage: big.NewFromString("0.01"),
			}

			assert.EqualValues(t, "-10.00", order.Profit().FormattedString(2))
		})

		t.Run("sell", func(t *testing.T) {
			order := Order{
				Side:          SELL,
				Amount:        big.NewFromString("1"),
				Price:         big.NewFromString("10"),
				FeePercentage: big.NewFromString("0.01"),
			}

			assert.EqualValues(t, "9.90", order.Profit().FormattedString(2))
		})
	})

	t.Run("without commission", func(t *testing.T) {
		t.Run("buy", func(t *testing.T) {
			order := Order{
				Side:          BUY,
				Amount:        big.NewFromString("1"),
				Price:         big.NewFromString("10"),
				FeePercentage: big.NewFromString("0"),
			}

			assert.EqualValues(t, "-10.00", order.Profit().FormattedString(2))
		})

		t.Run("sell", func(t *testing.T) {
			order := Order{
				Side:          SELL,
				Amount:        big.NewFromString("1"),
				Price:         big.NewFromString("10"),
				FeePercentage: big.NewFromString("0"),
			}

			assert.EqualValues(t, "10.00", order.Profit().FormattedString(2))
		})
	})
}
