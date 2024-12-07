package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	go BuildConn()
	go BuildConn()

	time.Sleep(time.Second * 30)
}

func BuildConn() {
	// 连接服务器
	conn, err := net.Dial("tcp", "localhost:8888")
	if err != nil {
		fmt.Println("Error connecting:", err.Error())
		return
	}
	defer conn.Close()

	// 发送消息
	message := "Hello from client!"
	conn.Write([]byte(message))

	// 接收响应
	buffer := make([]byte, 1024)
	_, err = conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
		return
	}

	fmt.Println("Received from server:", string(buffer))
}
