package mysql

import (
	"github.com/dystopia-systems/alaskalog"
	"os"
)
import "gorm.io/driver/mysql"
import "gorm.io/gorm"

var _db *gorm.DB

func InitDb() error {
	connString := os.Getenv("MYSQL_CONN_STRING")

	alaskalog.Logger.Infoln("Opening MySQL connection...")

	db, err := gorm.Open(mysql.Open(connString), &gorm.Config{})

	if err != nil {
		return err
	}

	_db = db

	alaskalog.Logger.Infoln("MySql connection successful.")

	return nil
}

func GetDb() *gorm.DB {
	return _db
}