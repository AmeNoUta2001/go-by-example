package main

import (
	"bufio"
	"log"
	"net"
)

/*
Socks5协议：让授权的用户可以通过单个端口访问内部的所有资源
运行此程序不能使用run命令
需要使用nc命令 nc可以直接与某个ip+端口进行tcp链接
*/
func main() {
	server, err := net.Listen("tcp", "127.0.0.1:1080")
	if err != nil {
		panic(err)
	}
	for {
		client, err := server.Accept()
		if err != nil {
			log.Printf("Accept failed %v", err)
			continue
		}
		go process(client)
	}
}

func process(conn net.Conn) {
	//退出函数时关闭链接
	defer conn.Close()

	reader := bufio.NewReader(conn)
	for {
		b, err := reader.ReadByte()
		if err != nil {
			break
		}
		_, err = conn.Write([]byte{b})
		if err != nil {
			break
		}
	}
}
