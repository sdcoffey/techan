package talib4g

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

// Decimal represents fixed precision numbers
type Decimal struct {
	unscaled int64
	scale    int
}

// NewDecimal creates a new decimal number equal to
// unscaled ** 10 ^ (-scale)
func NewDecimal(unscaled int64, scale int) Decimal {
	return Decimal{unscaled: unscaled, scale: scale}
}

func NewDecimalFromString(dec string) (d Decimal, err error) {
	err = d.UnmarshalText([]byte(dec))
	return
}

func NewDecimalFromFloat(dec float64) Decimal {
	b := []byte(fmt.Sprint(dec))
	d := new(Decimal)
	d.UnmarshalText(b)

	return *d
}

// MarshalText outputs a decimal representation of the scaled number
func (d Decimal) MarshalText() (text []byte, err error) {
	b := new(bytes.Buffer)
	if d.scale <= 0 {
		b.WriteString(strconv.FormatInt(d.unscaled, 10))
		b.WriteString(strings.Repeat("0", -d.scale))
	} else {
		str := strconv.FormatInt(d.unscaled, 10)
		if len(str) < d.scale {
			str = strings.Repeat("0", d.scale) + str
		}
		b.WriteString(str[:len(str)-d.scale])
		b.WriteString(".")
		b.WriteString(str[len(str)-d.scale:])
	}
	return b.Bytes(), nil
}

// UnmarshalText creates a Decimal from a string representation (e.g. 5.20)
// Currently only supports decimal strings
func (d *Decimal) UnmarshalText(text []byte) (err error) {
	var (
		str            = string(text)
		unscaled int64 = 0
		scale    int   = 0
	)

	if str == "" {
		return nil
	}

	if i := strings.Index(str, "."); i != -1 {
		scale = len(str) - i - 1
		str = strings.Replace(str, ".", "", 1)
	}

	if unscaled, err = strconv.ParseInt(str, 10, 64); err != nil {
		return err
	}

	d.unscaled = unscaled
	d.scale = scale

	return nil
}

func (x Decimal) cmp(y Decimal) int {
	xUnscaled, yUnscaled := x.unscaled, y.unscaled
	xScale, yScale := x.scale, y.scale

	for ; xScale > yScale; xScale-- {
		yUnscaled = yUnscaled * 10
	}

	for ; yScale > xScale; yScale-- {
		xUnscaled = xUnscaled * 10
	}

	switch {
	case xUnscaled < yUnscaled:
		return -1
	case xUnscaled > yUnscaled:
		return 1
	default:
		return 0
	}
}

func (d Decimal) Add(x Decimal) Decimal {
	maxScale, dRescaled, xRescaled := rescale(d, x)
	return Decimal{
		unscaled: dRescaled + xRescaled,
		scale:    maxScale,
	}
}

func (d Decimal) Sub(x Decimal) Decimal {
	maxScale, dRescaled, xRescaled := rescale(d, x)
	return Decimal{
		unscaled: dRescaled - xRescaled,
		scale:    maxScale,
	}
}

func (d Decimal) Mul(x Decimal) Decimal {
	maxScale, dRescaled, xRescaled := rescale(d, x)
	if maxRescaled := int64(Max(int(dRescaled), int(xRescaled))); maxRescaled == xRescaled {
		maxScale += numDigits(int(xRescaled)) - numDigits(int(dRescaled))
	} else {
		maxScale += numDigits(int(dRescaled)) - numDigits(int(xRescaled))
	}
	return Decimal{
		unscaled: dRescaled * xRescaled,
		scale:    maxScale,
	}
}

func numDigits(x int) (d int) {
	for x != 0 {
		x /= 10
		d++
	}

	return d
}

func rescale(x, y Decimal) (maxScale int, xRescaled, yRescaled int64) {
	maxScale = Max(x.scale, y.scale)
	xRescaled = x.unscaled
	yRescaled = y.unscaled

	if maxScale == x.scale {
		xRescaled *= int64(Pow(10, maxScale-y.scale))
	} else {
		yRescaled *= int64(Pow(10, maxScale-x.scale))
	}

	return
}

// String returns string representation of Decimal
func (d *Decimal) String() string {
	b, err := d.MarshalText()

	if err != nil {
		panic(err) //should never happen (see: MarshalText)
	}

	return string(b)
}
