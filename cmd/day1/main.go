package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var totalFuelNoExtra int = 0
	var totalFuel int = 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		mass, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}

		fuelNoExtra := CalcFuel(mass)
		totalFuelNoExtra += fuelNoExtra

		fuel := CalcFuelTotal(mass)
		totalFuel += fuel

		log.Printf("mass: %v, fuel (w/o extra): %v, fuel: %v\n", mass, fuelNoExtra, fuel)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	log.Printf("total fuel for all modules: %v\n", totalFuelNoExtra)
	log.Printf("total fuel for all modules and fuel: %v\n", totalFuel)

	// Note: incorrect method of calculating fuel:
	//     totalFuelNoExtra+CalcFuelTotal(totalFuelNoExtra)
}
