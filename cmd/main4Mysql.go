package main

import (
	"danilWeb/internal/dbs"
	"danilWeb/internal/user/dao"
	"danilWeb/internal/user/entity"
	"log"
)

func main4Mysql() {
	//http服务监听端口
	//service.StartUp()

	//数据库操作mysql
	MysqlDb := dbs.Init()
	defer MysqlDb.Close()

	//insert into
	//u := entity.User{
	//	Name: "老王",
	//	Age: 28,
	//}
	//err := dao.Add(u,MysqlDb)
	//log.Println("插入数据：是否错误",err)

	//find by id
	userByID := dao.StructQueryField(1, MysqlDb) //dbs.StructQueryField(1)
	println("+++StructQueryField++", userByID.Name)

	//find by name
	userByName := dao.FindByName("老王", MysqlDb)
	println("+++find by name++", userByName.Id, userByName.Name, userByName.Age)

	//update
	user_Update := dao.FindByName("老王1", MysqlDb)
	if nil != user_Update {
		user_Update.Name = "老王2"
		user_Update.Age = 22
		err := dao.Update(*user_Update, MysqlDb)
		if nil != err {
			log.Println("更新数据失败 : ", err)
		} else {
			log.Println("更新数据成功")
		}
	}

	//find all
	log.Println("******第一次全部查询")
	var users_findAll1 []entity.User = dao.FindAllUser(MysqlDb)
	for i := 0; i < len(users_findAll1); i++ {
		u := users_findAll1[i]
		println(u.Id, u.Name, u.Age)
	}

	//delete user
	user_Delete := dao.FindByName("老王2", MysqlDb)
	if nil != user_Delete {
		err := dao.Delete(*user_Delete, MysqlDb)
		if nil != err {
			log.Println("删除数据失败 : ", err)
		} else {
			log.Println("删除数据成功")
		}
	}

	//find all
	log.Println("******第二次全部查询")
	var aa2 []entity.User = dao.FindAllUser(MysqlDb)
	for i := 0; i < len(aa2); i++ {
		u := aa2[i]
		println(u.Id, u.Name, u.Age)
	}

}
