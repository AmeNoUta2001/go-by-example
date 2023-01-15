package main

import "fmt"

func main() {
	/*
		对于一个slice或者map可以使用range快速遍历
		使用range遍历数组时会返回两个值
		一个是索引 第二个是对应位置的值
	*/
	nums := []int{2, 3, 4}
	sum := 0
	for i, num := range nums {
		sum += num
		if num == 2 {
			fmt.Println("index:", i, "num:", num) // index: 0 num: 2
		}
	}
	fmt.Println(sum) // 9
	// 使用range遍历map时返回的第一个值是key，第二个值是value
	m := map[string]string{"a": "A", "b": "B"}
	for k, v := range m {
		fmt.Println(k, v) // b 8; a A
	}
	for k := range m {
		fmt.Println("key", k) // key a; key b
	}
}
