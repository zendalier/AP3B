package main

import (
	"fmt"
)

type car struct {
	brand      string
	model      string
	horsepower int
	weight     int
}

func (n car) Ratio() float64 {
	return float64(n.horsepower) / float64(n.weight)
}

func main() {
	supercar := [3]car{
		{
			brand:      "Lamborghini",
			model:      "Revuelto",
			horsepower: 1015,
			weight:     1772,
		},
		{
			brand:      "Ferrari",
			model:      "SF90 Stradale",
			horsepower: 986,
			weight:     1570,
		},
		{
			brand:      "McLaren",
			model:      "Artura",
			horsepower: 690,
			weight:     1395,
		},
	}

	for _, car := range supercar {
		fmt.Printf("Brand: %s\n", car.brand)
		fmt.Printf("Model: %s\n", car.model)
		fmt.Printf("Power to Weight Ratio: %.2f\n", car.Ratio())
		fmt.Println()
	}
}
