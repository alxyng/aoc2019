package main

import (
	"fmt"
	"testing"
)

var calcFuelData = []struct {
	mass int
	fuel int
}{
	{0, 0},
	{5, 0},
	{9, 1},
	{12, 2},
	{14, 2},
	{1969, 654},
	{100756, 33583},
}

func TestCalcFuel(t *testing.T) {
	for _, tt := range calcFuelData {
		t.Run(fmt.Sprintf("test mass %v", tt.mass), func(t *testing.T) {
			fuel := CalcFuel(tt.mass)
			if fuel != tt.fuel {
				t.Errorf("incorrect fuel, got %v want %v", fuel, tt.fuel)
			}
		})
	}
}

var calcFuelTotalData = []struct {
	mass int
	fuel int
}{
	{14, 2},
	{2, 0},
	{1969, 966},
	{654, 312},
	{100756, 50346},
	{33583, 16763},
}

func TestCalcFuelTotal(t *testing.T) {
	for _, tt := range calcFuelTotalData {
		t.Run(fmt.Sprintf("test mass %v", tt.mass), func(t *testing.T) {
			fuel := CalcFuelTotal(tt.mass)
			if fuel != tt.fuel {
				t.Errorf("incorrect fuel, got %v want %v", fuel, tt.fuel)
			}
		})
	}
}
