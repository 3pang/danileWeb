package main

import (
	"container/list"
	"fmt"
)

func main44() {
	s := "aa"
	b := "bb"
	c := fmt.Sprint(s, b)
	fmt.Println(c)
	d := new(list.List)
	d.PushFront(1)
	d.PushFront("s")
	fmt.Println(d)
	e := make(map[string]interface{})
	e["a"] = 1
	e["b"] = "2"
	delete(e, "a")

}
