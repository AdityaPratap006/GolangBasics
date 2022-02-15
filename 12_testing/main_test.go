package main

import (
	"math"
	"testing"
)

type Test struct {
	name     string
	dividend float64
	divisor  float64
	expected float64
	isErr    bool
}

var tests = []Test{
	{"valid-data", 100.0, 10.0, 10.0, false},
	{"invalid-data", 100.0, 0.0, 0.0, true},
	{"expect-5", 50.0, 10.0, 5.0, false},
	{"expect-fraction", -1.0, -777.0, 0.001287, false},
}

func TestDivision(t *testing.T) {
	for _, tt := range tests {
		got, err := divide(tt.dividend, tt.divisor)

		if tt.isErr {
			if err == nil {
				t.Error("Expected an error but did NOT get one")
			}
		} else {
			if err != nil {
				t.Error("Did NOT expect an error but got one:", err.Error())
			}
		}

		if !matchTo4DecimalPlaces(got, tt.expected) {
			t.Errorf("expected %f but got %f", roundDecimal(tt.expected, 4), roundDecimal(got, 4))
		}
	}
}

func matchTo4DecimalPlaces(x, y float64) bool {
	return roundDecimal(x, 4) == roundDecimal(y, 4)
}

func roundDecimal(number float64, places int) float64 {
	var pow10 = math.Pow10(places)
	return math.Floor(number*pow10) / pow10
}
