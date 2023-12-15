package product

import (
	"database/sql"
	"log"

	. "iknowbook.com/app/data"
	. "iknowbook.com/app/repository"

	"github.com/jmoiron/sqlx"
)

type ProductRepository struct {
	Connection *sqlx.DB
}

func NewProductRepository(conn *sqlx.DB) ProductRepository {
	return ProductRepository{
		Connection: conn,
	}
}

func (repo ProductRepository) queryEntity(sqlStr string, params ...interface{}) []Product {
	return QueryEntity[Product](repo.Connection, sqlStr, params...)
}

func (serv ProductRepository) QueryWithLimit(limit int) []Product {
	return serv.queryEntity(NewProductSqlStr().QueryWithLimit, limit)
}

func (serv ProductRepository) QueryWithPriceRange(min int, max int) []Product {
	return serv.queryEntity(NewProductSqlStr().QueryWithPriceRange, min, max)
}

func (serv ProductRepository) QueryByBarcode(code string) []Product {
	return serv.queryEntity(NewProductSqlStr().QueryByBarcode, code)
}

func (serv ProductRepository) execSql(str string, ps []Product) sql.Result {
	return ExecSql[[]Product](serv.Connection, str, ps)
}

func (serv ProductRepository) Insert(ps []Product) sql.Result {
	return serv.execSql(NewProductSqlStr().Insert, ps)
}

func (serv ProductRepository) Update(ps []Product) sql.Result {
	return serv.execSql(NewProductSqlStr().Update, ps)
}

func (serv ProductRepository) Delete(ps []Product) sql.Result {
	return serv.execSql(NewProductSqlStr().Delete, ps)
}

func (repo ProductRepository) QueryByConditions(
	pmin int,
	pmax int,
	p Product,
) []Product {
	return repo.queryEntity(
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
			log.Println(err)
		}
	}
	err = rows.Err()
	return max
}

func (repo ProductRepository) QueryNewest(ran int) []Product {
	return repo.queryEntity(NewProductSqlStr().QueryNewest, ran)
}
