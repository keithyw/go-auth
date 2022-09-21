package database

import (
	"log"
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"github.com/keithyw/auth/conf"
)

type MysqlDB struct {
	Config *conf.Config
	DB *sql.DB
}

func NewDatabase(config *conf.Config) (*MysqlDB) {
	mysqlConfig := mysql.Config{
		User: config.MysqlUser,
		Passwd: config.MysqlPass,
		Net: "tcp",
		Addr: config.MysqlHost,
		DBName: config.MysqlDBName,
		AllowNativePasswords: true,
	}
	db, err := sql.Open("mysql", mysqlConfig.FormatDSN())
	if err != nil {
		log.Printf("open failed %s", err.Error())
		panic(err)
	}
	if err = db.Ping(); err != nil {
		log.Printf("ping failed: %s", err.Error())
		panic(err)
	}
	log.Println("Connected to Mysql")
	return &MysqlDB{
		config, 
		db, 
	}
}