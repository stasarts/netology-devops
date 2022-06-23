package main

import "fmt"

func MetToFeet(met float64) (feet float64) {
	feet = met * 0.3048
	return
}

func main() {
	fmt.Print("Введите количество метров: ")
	var input float64
	fmt.Scanf("%f", &input)
	output := MetToFeet(input)
	fmt.Print("В ", input, " метрах: ", output, " футов.\n")
}
