package product

import (
	"database/sql"

	. "iknowbook.com/data"
	. "iknowbook.com/repository"

	"github.com/jmoiron/sqlx"
)

type ProductRepository struct {
	Connection *sqlx.DB
}

func NewProductRepository(conn *sqlx.DB) ProductRepository {
	return ProductRepository{Connection: conn}
}

func (serv ProductRepository) QueryEntity(sqlStr string, params ...interface{}) []Product {
	return QueryEntity[Product](serv.Connection, sqlStr, params...)
}

func (serv ProductRepository) QueryWithLimit(limit int) []Product {
	return serv.QueryEntity(NewProductSqlStr().QueryWithLimit, limit)
}

func (serv ProductRepository) QueryWithPriceRange(min int, max int) []Product {
	return serv.QueryEntity(NewProductSqlStr().QueryWithPriceRange, min, max)
}

func (serv ProductRepository) QueryByBarcode(code string) []Product {
	return serv.QueryEntity(NewProductSqlStr().QueryByBarcode, code)
}

func (serv ProductRepository) ExecSql(str string, ps []Product) sql.Result {
	return ExecSql[[]Product](serv.Connection, str, ps)
}

func (serv ProductRepository) Insert(ps []Product) sql.Result {
	return serv.ExecSql(NewProductSqlStr().Insert, ps)
}

func (serv ProductRepository) Update(ps []Product) sql.Result {
	return serv.ExecSql(NewProductSqlStr().Update, ps)
}

func (serv ProductRepository) Delete(ps []Product) sql.Result {
	return serv.ExecSql(NewProductSqlStr().Delete, ps)
}

func (serv ProductRepository) QueryByConditions(
	pmin int,
	pmax int,
	p Product,
) []Product {
	return QueryEntity[Product](
		serv.Connection,
		NewProductSqlStr().QueryByConditions,
		pmin,
		pmax,
		p.Product_title,
		p.Publisher,
	)
}

func (repo ProductRepository) MaxPrice() int {
	rows, err := repo.Connection.Query(NewProductSqlStr().MaxPrice)
	var max int
	for rows.Next() {
		err = rows.Scan(&max)
		if err != nil {
			panic(err)
		}
	}
	err = rows.Err()
	return max
}
