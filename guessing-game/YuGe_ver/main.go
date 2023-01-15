package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixMicro())
	RandomNumber := rand.Intn(100) + 1
	fmt.Println(RandomNumber)
	var guess int
	// for i := 0; i < 20; i++ {
	// 	fmt.Println(RandomNumber)
	// }
	for {
		fmt.Println("请输入一个1-100之间的整数")
		fmt.Scanf("%d", &guess)
		if guess > RandomNumber {
			fmt.Println("猜大了")
		} else if guess < RandomNumber {
			fmt.Println("猜小了")
		} else {
			fmt.Println("猜对了！")
			break
		}
	}
}
