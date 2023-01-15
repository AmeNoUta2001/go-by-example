package main

import "fmt"

func main() {
	/*
		make关键字也可以创建map（java中叫HashMap）
	*/
	m := make(map[string]int)
	m["one"] = 1
	m["two"] = 2
	fmt.Println(m)           // map[one:1 two:2]
	fmt.Println(len(m))      // 2
	fmt.Println(m["one"])    // 1
	fmt.Println(m["unknow"]) // 0

	// 这一句中的ok用来获取在map中这个key是否存在
	r, ok := m["unknow"]
	fmt.Println(r, ok) // 0 false

	delete(m, "one")
	// golang中的map顺序完全随机 在输出是也会按照随机的顺序输出元素
	m2 := map[string]int{"one": 1, "two": 2}
	var m3 = map[string]int{"one": 1, "two": 2}
	fmt.Println(m2, m3)
}
