package dao

import (
	"danilWeb/internal/user/entity"
	"database/sql"
	"fmt"
	"log"
)

// 根据ID查询对象
func StructQueryField(userId int64, MysqlDb *sql.DB) *entity.User {
	user := new(entity.User)
	log.Println("##########start row")
	row, _ := MysqlDb.Query("SELECT u.id, u.name, u.age FROM users u WHERE u.id=?", userId)

	if row.Next() {
		log.Println("##########start scab")
		if err := row.Scan(&user.Id, &user.Name, &user.Age); err != nil {
			log.Println("*****************scan is nil")
			log.Fatal("scan failed, err:%v", err)
			fmt.Print(err)
		}
		log.Println("user1========")
		log.Println(user.Id, user.Name, user.Age)
	}
	return user
}

// 根据Name查询对象
func FindByName(userName string, MysqlDb *sql.DB) *entity.User {
	user := new(entity.User)
	log.Println("******dao.FindByName start row")
	row, _ := MysqlDb.Query("SELECT id, name, age FROM users WHERE name =?", userName)

	if row.Next() {
		log.Println("******dao.FindByName start scab")
		if err := row.Scan(&user.Id, &user.Name, &user.Age); err != nil {
			log.Println("******dao.FindByName scan is nil")
			log.Fatal("scan failed, err:%v", err)
			fmt.Print(err)
		}
		log.Println("******dao.FindByName user1=")
		log.Println(user.Id, user.Name, user.Age)
	}
	return user
}

//查询所有对象
func FindAllUser(MysqlDb *sql.DB) []entity.User {
	users := make([]entity.User, 0)
	row, _ := MysqlDb.Query("SELECT u.id, u.name, u.age FROM users u")
	for row.Next() {
		user := entity.User{}
		row.Scan(&user.Id, &user.Name, &user.Age)
		users = append(users, user)
	}
	return users
}

//插入数据
func Add(userpara entity.User, MysqlDb *sql.DB) error {
	ret, err := MysqlDb.Exec("INSERT INTO users(name,age) VALUES (?,?)", userpara.Name, userpara.Age)
	if nil == err {
		//插入数据的主键id
		lastInsertID, _ := ret.LastInsertId()
		log.Println("*****dao.Add LastInsertID = ", lastInsertID)
		rowaffectid, _ := ret.RowsAffected() //影响行数
		log.Println("*****dao.Add Rowaffectid = ", rowaffectid)
	}
	return err
}

//更新
func Update(userpara entity.User, MysqlDb *sql.DB) error {
	ret, err := MysqlDb.Exec("UPDATE users set name = ?,age = ? WHERE id = ?", userpara.Name, userpara.Age, userpara.Id)
	if nil != err {
		log.Println("dao.update 更新失败：", err)
	} else {
		rowsAffected_nums, _ := ret.RowsAffected() //影响行数
		log.Println("******dao.update rowsAffected_nums :", rowsAffected_nums)
	}
	return err
}

//删除
func Delete(userpara entity.User, MysqlDb *sql.DB) error {
	ret, err := MysqlDb.Exec("DELETE FROM users WHERE id = ?", userpara.Id)
	if nil != err {
		log.Println("dao.delete 删除失败：", err)
	} else {
		rowsAffected_nums, _ := ret.RowsAffected() //影响行数
		log.Println("******dao.delete rowsAffected_nums :", rowsAffected_nums)
	}
	return err
}
