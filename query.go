package ddorm

import (
	"database/sql"
)

// Find all records matching model criteria
func (d *DDORM) Find(model interface{}) (results []interface{}, err error) {
	tableName := GetTableName(model)

	var rows *sql.Rows
	rows, err = d.DB.Query("SELECT * FROM " + tableName)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	// Needs fixing
	// results, err = GetResultsFromRows(rows, model)

	return results, nil
}
