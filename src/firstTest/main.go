package main

import (
	"testing"
	"github.com/StivenP6160/tester/testunitario"
)

func TestSuma(t *testing.T) {
	total := Suma(5, 6)

	if total != 10 {
		t.Errorf("Suma incorrecta, tiene %d se esperaba %d", total, 10)
	}
}