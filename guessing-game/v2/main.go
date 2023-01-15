package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	maxNum := 100
	/*
		在使用rand时要设置随机数种子
		否则每次会生成同样的随机数序列
	*/
	rand.Seed(time.Now().UnixNano())
	secretNumber := rand.Intn(maxNum)
	fmt.Println("The secret number is ", secretNumber)
}
