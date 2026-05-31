package main

import "testing"

func TestStandardPricing(t *testing.T) {
	tests := []struct {
		basePrice	float64
		nights		int
		expected	float64
	}{
		{100, 2, 200},
		{50, 3, 150},
		{125, 4, 500},
		{77.35, 5, 386.75},
	}
	for _, tt := range tests {
		p := StandardPricing{}
		result := p.Calculate(tt.basePrice, tt.nights)
		if result != tt.expected {
			t.Errorf("StandardPricing(%f, %d) = %f, want %f", tt.basePrice, tt.nights, result, tt.expected)
		}
	}
}

func TestWeekendPricing(t *testing.T) {
	tests := []struct {
		basePrice	float64
		nights		int
		expected	float64
	}{
		{100, 2, 300},
		{50, 3, 225},
		{125, 4, 750},
		{77.35, 5, 580.125},
	}
	for _, tt := range tests {
		p := WeekendPricing{}
		result := p.Calculate(tt.basePrice, tt.nights)
		if result != tt.expected {
			t.Errorf("WeekendPricing(%f, %d) = %f, want %f", tt.basePrice, tt.nights, result, tt.expected)
		}
	}
}