package system

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type DBConfig struct {
	Name     string
	Password string
	Host     string
	Port     string
	DBName   string
}

func InitDatabaseConnection(dataSourceName DBConfig) (*sql.DB, error) {
	// Open the connection with MySQL DB.
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dataSourceName.Name, dataSourceName.Password, dataSourceName.Host, dataSourceName.Port, dataSourceName.DBName))
	if err != nil {
		return nil, err
	}

	// Check the connection.
	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, err
}
