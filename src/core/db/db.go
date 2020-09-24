package db

import "os"
import "gorm.io/driver/mysql"
import "gorm.io/gorm"

var _db *gorm.DB

func InitDb() error {
	connString := os.Getenv("MY_SQL_GORM_CONN_STRING")

	db, err := gorm.Open(mysql.Open(connString), &gorm.Config{})

	if err != nil {
		return err
	}

	_db = db

	return nil
}

func GetDb() *gorm.DB {
	return _db
}