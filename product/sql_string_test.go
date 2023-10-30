package product

import "testing"

func TestNewProductSqlStr(t *testing.T) {
	sqlS := NewProductSqlStr()
	t.Log(sqlS.QueryWithLimit)
	t.Log(sqlS.QueryWithPriceRange)
	t.Log(sqlS.QueryByBarcode)
	t.Log(sqlS.QueryByConditions)
	t.Log(sqlS.QueryNewest)

	t.Log(sqlS.Insert)
	t.Log(sqlS.Delete)
	t.Log(sqlS.Update)
}
