package main

import "math"

// CalcFuelTotal calculates the fuel required to launch a module with a given
// mass, taking into consideration the mass of the added fuel.
func CalcFuelTotal(mass int) int {
	fuel := CalcFuel(mass)
	if fuel == 0 {
		return 0
	}

	return fuel + CalcFuelTotal(fuel)
}

// CalcFuel calculates the fuel required to launch a module with a given mass.
func CalcFuel(mass int) int {
	return maxInt((mass/3)-2, 0)
}

func maxInt(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}
