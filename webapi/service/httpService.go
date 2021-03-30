package service

import (
	"danilWeb/webapi/loginapi"
	"log"
	"net/http"
)

func StartUp() {
	h := http.FileServer(http.Dir("website"))                  //目录位置
	h2 := http.FileServer(http.Dir("htmltext"))                //目录位置
	http.Handle("/website/", http.StripPrefix("/website/", h)) //启动静态文件服务
	http.Handle("/ht/", http.StripPrefix("/ht/", h2))          //启动静态文件服务，url访问的路径
	http.HandleFunc("/login", loginapi.Login)                  //设置访问路由
	http.HandleFunc("/login2", loginapi.Login2)                //设置访问路由
	err := http.ListenAndServe(":9077", nil)                   //端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	} else {
		log.Println("=========ListenAndServe is ready")
	}
}
