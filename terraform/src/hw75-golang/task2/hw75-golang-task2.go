package main

import "fmt"

func Min(listToFindMin []int) (min int) {
	current := 0
	for i, next := range listToFindMin {
		if i == 0 {
			current = next
		} else {
			if next < current {
				current = next
			}
		}
	}
	return current
}

func main() {
	x := []int{48, 96, 86, 68, 57, 82, 63, 70, 37, 34, 83, 27, 19, 97, 9, 17}
	fmt.Println("Заданный список: ", x)
	y := Min(x)
	fmt.Print("Наименьший элемент в заданном списке: ", y, ".\n")
}
