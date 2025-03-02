package net_code

import (
	"fmt"
	"net"
	"os"
)

func TCPServer() {
	listener, err := net.Listen("tcp", "localhost:8089")
	if err != nil {
		fmt.Println("error starting server: ", err)
		os.Exit(1)
	}
	defer listener.Close()

	fmt.Println("Server listening on localhost:8089")

	for {
		//等待连接
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("error accept connetction: ", err)
			continue
		}
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	buffer := make([]byte, 1024)
	_, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("error Reading: ", err)
	}
	conn.Write(buffer)
	conn.Close()
}

func TCPClient() {
	//连接服务器
	conn, err := net.Dial("tcp", "localhost:8089")
	if err != nil {
		fmt.Println("error connetion: ", err)
		os.Exit(1)
	}
	defer conn.Close()

	//发送数据
	_, err = conn.Write([]byte("hello server!"))
	if err != nil {
		fmt.Println("error write: ", err)
	}
	//读取响应
	buffer := make([]byte, 1024)
	_, err = conn.Read(buffer)
	if err != nil {
		fmt.Println("error reading: ", err)
	}
	fmt.Println(string(buffer))
}
