package techan

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommidityChannelIndexIndicator_Calculate(t *testing.T) {
	typicalPrices := []string{
		"23.98", "23.92", "23.79", "23.67", "23.54",
		"23.36", "23.65", "23.72", "24.16", "23.91",
		"23.81", "23.92", "23.74", "24.68", "24.94",
		"24.93", "25.10", "25.12", "25.20", "25.06",
		"24.50", "24.31", "24.57", "24.62", "24.49",
		"24.37", "24.41", "24.35", "23.75", "24.09",
	}

	series := mockTimeSeries(typicalPrices...)

	cci := NewCCIIndicator(series, 20)

	results := []string{"101.9185", "31.1946", "6.5578", "33.6078", "34.9686", "13.6027",
		"-10.6789", "-11.4710", "-29.2567", "-128.6000", "-72.7273"}

	for i, result := range results {
		assert.EqualValues(t, result, cci.Calculate(i+19).FormattedString(4))
	}
}
