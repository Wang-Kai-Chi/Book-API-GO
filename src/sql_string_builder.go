package main

import (
	"reflect"
	"strconv"
	"strings"
)

func GetStructFieldNames(elm reflect.Value) []string {
	var out []string

	for i := 0; i < elm.NumField(); i++ {
		n := elm.Type().Field(i).Name
		out = append(out, strings.ToLower(n))
	}
	return out
}

type SqlBuild[T AData] interface {
	GetInsertSQLString() string
}

type SqlStringBuilder[T AData] struct {
	Data      []T
	Form      T
	TableName string
	Ids       []string
}

func RemoveStr(strs []string, i int) []string {
	copy(strs[i:], strs[i+1:])
	strs[len(strs)-1] = ""
	strs = strs[:len(strs)-1]
	return strs
}
func (builder SqlStringBuilder[AData]) GetInsertSQLString() string {
	form := builder.Form
	elm := reflect.ValueOf(&form).Elem()

	table := builder.TableName
	sqlStr := "INSERT INTO " + table + "("

	getFieldsWithoutIds := func(fields []string, ids []string) []string {
		for i := 0; i < len(fields); i++ {
			for j := 0; j < len(ids); j++ {
				if fields[i] == strings.ToLower(ids[j]) {
					fields = RemoveStr(fields, i)
				}
			}
		}
		return fields
	}

	ids := builder.Ids
	fields := getFieldsWithoutIds(GetStructFieldNames(elm), ids)
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
	data := builder.Data
	for i := 0; i < len(data); i++ {
		inserts = append(inserts, paramString)
	}
	insertVals := strings.Join(inserts, ",")
	sqlStr = sqlStr + insertVals

	return sqlStr
}
