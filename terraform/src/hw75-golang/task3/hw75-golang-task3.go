package main

import "fmt"

func main() {
	fmt.Println("Числа от 1 до 100, которые делятся на 3: ")
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()
}
