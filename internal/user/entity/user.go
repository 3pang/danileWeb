package entity

// 用户表结构体
type User struct {
	Id   int64  `db:"id"`
	Name string `db:"name"`
	Age  int    `db:"age"`
}
