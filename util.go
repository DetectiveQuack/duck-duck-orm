package ddorm

import (
	"database/sql"
	"reflect"
)

// GetTableName gets table name from model
func GetTableName(model interface{}) string {
	return reflect.TypeOf(model).Elem().Name()
}

// GetNewInstance gets a new blank instance of a model
func GetNewInstance(model interface{}) interface{} {
	return reflect.New(reflect.TypeOf(model))
}

// GetResultsFromRows scans the results and returns an array of models
func GetResultsFromRows(rows *sql.Rows, model interface{}) (results []interface{}, err error) {
	for rows.Next() {
		m := GetNewInstance(model)
		err := rows.Scan(m)

		if err != nil {
			return nil, err
		}

		results = append(results, m)
	}

	return results, err
}
