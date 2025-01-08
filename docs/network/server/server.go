package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	// 监听端口
	listener, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}
	defer listener.Close()

	fmt.Println("Listening on localhost:8888")

	for {
		// 接受连接
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			continue
		}

		// 处理连接
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)
	_, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
		return
	}

	fmt.Println("Received:", string(buffer))
	fmt.Println("正在处理中")
	time.Sleep(10 * time.Second)

	// 发送响应
	response := "Hello from server!"
	conn.Write([]byte(response))
}
