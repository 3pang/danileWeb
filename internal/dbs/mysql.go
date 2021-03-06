package dbs

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

var MysqlDb *sql.DB
var MysqlDbErr error

const (
	USER_NAME = "root"
	//PASS_WORD = "123456"
	PASS_WORD = "Aa@123456"
	//HOST      = "localhost"
	HOST = "192.168.115.210"
	PORT = "3306"
	//DATABASE  = "test"
	DATABASE = "mysql223"
	CHARSET  = "utf8"
)

//init connection
func Init() *sql.DB {
	dbDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", USER_NAME, PASS_WORD, HOST, PORT, DATABASE, CHARSET)
	//open connection
	MysqlDb, MysqlDbErr := sql.Open("mysql", dbDSN)
	//defer MysqlDb.Close()
	if MysqlDbErr != nil {
		log.Println("This is a wrong dbDSN: " + dbDSN)
		panic("数据源配置不正确: " + MysqlDbErr.Error())
	} else {
		log.Println(" dbDSN is correct :" + dbDSN)
	}
	MysqlDb.SetMaxOpenConns(100)
	MysqlDb.SetMaxIdleConns(20)
	MysqlDb.SetConnMaxIdleTime(100 * time.Second)
	fmt.Println("----------MysqlDb.Ping():%s", MysqlDb.Ping())
	if MysqlDbErr = MysqlDb.Ping(); nil != MysqlDbErr {
		panic("数据库链接失败: " + MysqlDbErr.Error())
	} else {
		log.Println("==========db connect ready")
	}
	return MysqlDb
}
