package main

import "testing"

func TestNewProductSqlC(t *testing.T) {
	p := NewProductSqlC()
	t.Log(p.InsertSQL)
}

func TestProductSqlCInsert(t *testing.T) {
	sql := NewProductSqlC().GetInsertString()
	t.Log(sql)
}
