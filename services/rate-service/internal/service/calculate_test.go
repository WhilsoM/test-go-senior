package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateAvg(t *testing.T) {
	tests := []struct {
		name     string
		data     []float64
		n, m     int
		expected float64
	}{
		{
			name:     "valid range",
			data:     []float64{10.0, 20.0, 30.0, 40.0},
			n:        1,
			m:        2,
			expected: 25.0,
		},
		{
			name:     "out of range m",
			data:     []float64{10.0, 20.0},
			n:        0,
			m:        5,
			expected: 0,
		},
		{
			name:     "empty data",
			data:     []float64{},
			n:        0,
			m:        0,
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calculateAvg(tt.data, tt.n, tt.m)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestCalculateTopN(t *testing.T) {
	tests := []struct {
		name     string
		data     []float64
		n        int
		expected float64
	}{
		{
			name:     "valid index",
			data:     []float64{10.5, 20.1, 30.8},
			n:        1,
			expected: 20.1,
		},
		{
			name:     "index out of upper bound",
			data:     []float64{10.5, 20.1},
			n:        5,
			expected: 0,
		},
		{
			name:     "negative index",
			data:     []float64{10.5, 20.1},
			n:        -1,
			expected: 0,
		},
		{
			name:     "empty data",
			data:     []float64{},
			n:        0,
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calculateTopN(tt.data, tt.n)
			assert.Equal(t, tt.expected, result)
		})
	}
}
