package main

import (
	"danilWeb/internal/dbs"
	"danilWeb/internal/user/dao"
)

func main() {

	//service.StartUp()
	MysqlDb := dbs.Init()
	dao.StructQueryField(1, MysqlDb)
	MysqlDb.Close()
	//dbs.StructQueryField(1)
}
