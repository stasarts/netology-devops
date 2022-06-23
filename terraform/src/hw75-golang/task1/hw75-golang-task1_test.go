package main

import "testing"

func TestMetToFeet(t *testing.T) {
	var v float64 = MetToFeet(100)
	if v != 30.48 {
		t.Error("В 100 метрах 30.48 футов, а получилось ", v)
	}
}
