package databasefactory

import (
	"errors"
	"os"
	"tmp-report-transactions/repository/databasefactory/drivers"
)

func GetDatabase() (db Database, err error) {

	driverName := os.Getenv("app.database_driver")
	// Default driver will return MYSQL-Driver
	if driverName == "" {
		driverName = drivers.MYSQL
	}

	if driverName == drivers.MYSQL {
		return newMysqlImpl(), nil
	} else {
		return db, errors.New("unimplement database driver")
	}
}
