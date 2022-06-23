package main

import "testing"

func TestMin(t *testing.T) {
	var v int = Min([]int{48, 96, 86, 68, 57, 82, 63, 70, 37, 34, 83, 27, 19, 97, 9, 17})
	if v != 9 {
		t.Error("Минимальное число в списке - 9, а получилось ", v)
	}
}
