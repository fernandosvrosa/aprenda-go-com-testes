package main

import "testing"

func TestOla(t *testing.T) {
	resultado := Ola("Fernando")
	esperado := "Olá, Fernando"

	if resultado != esperado {
		t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
	}
}
