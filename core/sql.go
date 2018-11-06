package core

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func GetSqlConn() (*gorm.DB, error) {
	db := Conf.Db
	args := ""

	switch db.Driver {
	case "mysql":
		args = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=UTC",
			db.Username, db.Password, db.Address, db.Port, db.Dbname)
	case "postgres":
		args = fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s",
			db.Address, db.Port, db.Dbname, db.Username, db.Password)
	}

	return gorm.Open(db.Driver, args)
}