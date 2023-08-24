package main

import (
	"reflect"
	"testing"
)

func TestGetInsertSQLString(t *testing.T) {
	var p Product
	builder := SqlStringBuilder[Product]{
		Data:      getProductForTest(),
		Form:      p,
		TableName: "product",
		Ids:       []string{"Product_id"},
	}
	s := builder.GetInsertSQLString()
	t.Log(s)
}

func TestGetStructFieldNames(t *testing.T) {
	var p Product
	t.Log("product")
	t.Log(GetStructFieldNames(reflect.ValueOf(&p).Elem()))
}
