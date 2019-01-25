package db

import (

	// Mysql driver. Ignore this include
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

// Connect - creates new db engine
func Connect(database, host, port, user, password string) (db *xorm.Engine, err error) {
	db, err = xorm.NewEngine("mysql", user+":"+password+"@tcp("+host+":"+port+")/"+database+"?parseTime=true")
	if err != nil {
		return
	}

	db.SetConnMaxLifetime(1000)
	_, err = db.DBMetas()
	return
}
