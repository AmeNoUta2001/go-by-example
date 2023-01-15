package main

import "fmt"

func main() {

	// 在go语言中 数组的长度是固定的 所以在开发中一般不用
	var a [5]int
	a[4] = 100
	fmt.Println("get:", a[2])
	fmt.Println("len:", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
