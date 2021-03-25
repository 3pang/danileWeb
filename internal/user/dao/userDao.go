package dao

import (
	"danilWeb/internal/user/entity"
	"database/sql"
	"fmt"
	"log"
)

// 查询数据，指定字段名
func StructQueryField(userId int64, MysqlDb *sql.DB) {
	user := new(entity.User)
	log.Println("##########start row")
	row, _ := MysqlDb.Query("select id, name, age from users where id=?", 1)
	row.Next()
	log.Println("##########start scab")
	if err := row.Scan(&user.Id, &user.Name, &user.Age); err != nil {
		log.Println("*****************scan is nil")
		log.Fatal("scan failed, err:%v", err)
		fmt.Print(err)
		return
	}
	log.Println(user.Id, user.Name, user.Age)
}
