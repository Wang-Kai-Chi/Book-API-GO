package main

import (
	"context"
	"database/sql"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func (p Product) QueryAll(db *sql.DB, limit int64) ([]Product, error) {
	elem := reflect.ValueOf(&p).Elem()

	rows, err := db.Query(
		"SELECT " + strings.Join(getStructFields(elem), ",") + " FROM product LIMIT " + strconv.FormatInt(limit, 10))

	var products []Product

	for rows.Next() {
		var p Product

		err := rows.Scan(
			&p.Product_id,
			&p.Barcode,
			&p.Product_title,
			&p.Publisher,
			&p.Publication_date,
			&p.Price,
			&p.Quantity,
			&p.Description,
		)
		if err != nil {
			panic(err)
		}
		products = append(products, p)
	}
	db.Close()

	return products, err
}

func (p Product) Insert(db *sql.DB, ps []Product) (int64, error) {
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	sqlStr := func() string {
		sqlStr := "INSERT INTO product(barcode, product_title, publisher, publication_date, price, quantity, description) VALUES "

		var inserts []string
		for i := 0; i < len(ps); i++ {
			inserts = append(inserts, "($1, $2, $3, $4, $5, $6, $7)")
		}
		insertVals := strings.Join(inserts, ",")
		sqlStr = sqlStr + insertVals
		return sqlStr
	}
	stmt, err := db.PrepareContext(ctx, sqlStr())

	if err != nil {
		panic(err)
	}

	params := func() []interface{} {
		var params []interface{}
		for _, v := range ps {
			price, err := strconv.Atoi(strings.ReplaceAll(v.Price, "元", ""))

			if err == nil {
				params = append(params, v.Barcode, v.Product_title, v.Publisher, v.Publication_date, price, v.Quantity, v.Description)
			}
		}
		return params
	}
	res, err := stmt.ExecContext(ctx, params()...)
	if err != nil {
		panic(err)
	}
	rows, err := res.RowsAffected()

	return rows, err
}

func getStructFields(elm reflect.Value) []string {
	var out []string

	for i := 0; i < elm.NumField(); i++ {
		n := elm.Type().Field(i).Name
		out = append(out, strings.ToLower(n))
	}
	return out
}
func GetInsertSQLString[T AData](data []T, form T, table string, ids []string) string {
	elm := reflect.ValueOf(&form).Elem()

	sqlStr := "INSERT INTO " + table + "("

	getFieldsWithoutIds := func(fields []string, ids []string) []string {
		removeElement := func(fies []string, i int) []string {
			copy(fies[i:], fies[i+1:])
			fies[len(fies)-1] = ""
			fies = fies[:len(fies)-1]
			return fies
		}
		for i := 0; i < len(fields); i++ {
			for j := 0; j < len(ids); j++ {
				if fields[i] == strings.ToLower(ids[j]) {
					fields = removeElement(fields, i)
				}
			}
		}
		return fields
	}

	fields := getFieldsWithoutIds(getStructFields(elm), ids)
	sqlStr = sqlStr + strings.Join(fields, ",") + ") VALUES "

	getParamStr := func(fields []string) string {
		var params []string
		for i := 0; i < len(fields); i++ {
			temp := i + 1
			n := "$" + strconv.FormatInt(int64(temp), 10)
			params = append(params, n)
		}
		return "(" + strings.Join(params, ",") + ")"
	}

	paramString := getParamStr(fields)

	var inserts []string
	for i := 0; i < len(data); i++ {
		inserts = append(inserts, paramString)
	}
	insertVals := strings.Join(inserts, ",")
	sqlStr = sqlStr + insertVals

	return sqlStr
}
