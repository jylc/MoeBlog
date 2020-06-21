package conn

import (
	"MoeBlog/conf"
	"github.com/jinzhu/gorm"
	"log"
)

type mySqlDB struct {
	db *gorm.DB
}

var mysqlDb *mySqlDB

func InitMySql() {
	dsn := conf.MySQLConfig.Config.FormatDSN()
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("连接数据库失败：%s", err.Error())
	}
	mysqlDb = &mySqlDB{}
	mysqlDb.db = db
	mysqlDb.db.DB().SetConnMaxLifetime(conf.MySQLConfig.ConnMaxLifeTime)
	mysqlDb.db.DB().SetMaxIdleConns(conf.MySQLConfig.MaxIdleConns)
	mysqlDb.db.DB().SetMaxOpenConns(conf.MySQLConfig.MaxOpenConns)

	// 禁用默认表名的复数形式，如果置为 true，则 `User` 的默认表名是 `auth`
	mysqlDb.db.SingularTable(true)
	err = mysqlDb.db.DB().Ping()
	if err != nil {
		log.Fatalf("Ping数据库失败：%s", err.Error())
	}

	// 开启debug模式，打印sql语句
	db.LogMode(true)
}

func GetInstance() *gorm.DB {
	return mysqlDb.db
}
