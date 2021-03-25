package dbs

import (
	"danilWeb/internal/user/entity"
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
	PASS_WORD = "123456"
	HOST      = "localhost"
	PORT      = "3306"
	DATABASE  = "test"
	CHARSET   = "utf8"
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

// 查询数据，指定字段名
func StructQueryField(userId int64) {
	user := new(entity.User)
	log.Println("##########start row")
	row, _ := MysqlDb.Query("select id, name, age from users where id=?", 1)
	log.Println("##########start scab")
	if err := row.Scan(&user.Id, &user.Name, &user.Age); err != nil {
		log.Println("*****************scan is nil")
		log.Fatal("scan failed, err:%v", err)
		fmt.Print(err)
		return
	}
	log.Println(user.Id, user.Name, user.Age)
}
