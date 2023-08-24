package main

import "testing"

func TestNewProductSqlC(t *testing.T) {
	p := NewProductSqlC()
	t.Log(p.Insert)
}
