package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	_ "strings"
)

func login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println("method:", r.Method) //获取请求方法
	if r.Method == "GET" {
		log.Println("GET 方法进入")
		t, _ := template.ParseFiles("./website/helloworld.html")
		log.Println(t.Execute(w, nil))
	} else {
		log.Println("POST 方法进入")
		defer r.Body.Close()
		con, _ := ioutil.ReadAll(r.Body)
		fmt.Println(string(con))

		var dat map[string]interface{}
		if err := json.Unmarshal([]byte(string(con)), &dat); err == nil {
			fmt.Println(dat)
			fmt.Println(dat["username"])
		} else {
			fmt.Println(err)
		}
		/*
			//请求的是登录数据，那么执行登录的逻辑判断
			fmt.Println("username:", r.Form["username"])
			fmt.Println("password:", r.Form["password"])*/
	}
}
func main() {
	h := http.FileServer(http.Dir("website"))
	http.Handle("/website/", http.StripPrefix("/website/", h)) //启动静态文件服务
	http.HandleFunc("/login", login)                           //设置访问路由
	err := http.ListenAndServe(":9077", nil)                   //端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
