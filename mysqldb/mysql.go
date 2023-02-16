package mysqldb

import "github.com/jinzhu/gorm"

type MyDatabase struct {
	Db *gorm.DB
}

func (mc *MyDatabase) CloseMySQLDatabase() {
	mc.Db.Close()
}

func (mc *MyDatabase) Find(out interface{}, where ...interface{}) error {
	return mc.Db.Find(out, where...).Error
}
