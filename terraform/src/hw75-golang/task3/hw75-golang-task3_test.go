package main

import "testing"

func TestMain(t *testing.T) {
	var v []int
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			v = append(v, i)
		}
	}
	if v[0] != 3 || v[8] != 27 || v[15] != 48 || v[32] != 99 {
		t.Error("Ожидалось 3, 27, 48, 99, а получили", v[0], v[8], v[15], v[32])
	}
}
