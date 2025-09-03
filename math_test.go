package main

import "testing"

func TestSoma(t *testing.T) {
	total := soma(15, 14)
	if total != 29 {
		t.Errorf("Total %d, esperado %d", total, 29)
	}
}