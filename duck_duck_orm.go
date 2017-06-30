package ddorm

import (
	"database/sql"
)

// DBType type of database to connect to
type DBType string

// Holds the database connection names of type DBType
const (
	POSTGRES DBType = "postgres"
)

// DDORM struct holds database information
type DDORM struct {
	DB *sql.DB
}

// NewDDORM creates new ddorm instance
func NewDDORM() DDORM {
	d := DDORM{}

	return d
}

// Connect opens a connection to the database and pings it
func (d *DDORM) Connect(dbType DBType, dataSourceName string) (err error) {
	strDb := string(dbType)

	d.DB, err = sql.Open(strDb, dataSourceName)

	if err = d.DB.Ping(); err != nil {
		return err
	}

	return err
}
