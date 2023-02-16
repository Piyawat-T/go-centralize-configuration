package bootstrap

import (
	"fmt"

	"github.com/Piyawat-T/go-centralize-configuration/mysqldb"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Database interface {
	CloseMySQLDatabase()
	Find(out interface{}, where ...interface{}) error
}

func NewMySQLDatabase(env *Env) Database {
	dbHost := env.DBHost
	dbPort := env.DBPort
	dbUser := env.DBUser
	dbPass := env.DBPass
	dbName := env.DBName

	mysqlURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)

	db, err := gorm.Open("mysql", mysqlURI)
	if err != nil {
		fmt.Println("Status:", err)
	}
	return &mysqldb.MyDatabase{Db: db}
}
