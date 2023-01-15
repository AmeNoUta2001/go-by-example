package main

import "fmt"

/*
	在golang中 指针的操作非常有限
	指针的主要用处是对传入的参数进行修改
*/

/*
	该函数并不会修改传入的参数
	它修改的只是这个参数的拷贝
*/
func add2(n int) {
	n += 2
}

func add2ptr(n *int) {
	*n += 2
}

func main() {
	n := 5
	add2(n)
	fmt.Println(n) // 5
	add2ptr(&n)
	fmt.Println(n) // 7
}
