package main

import (
	"fmt"
	"net"
	"strings"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err)
		return
	}

	clientData := string(buffer[:n])
	clientData = strings.TrimSpace(clientData)

	// ตรวจสอบ username , password (กำหนดไว้ตามฌจท)
	if clientData == "std1:p@ssw0rd" {
		conn.Write([]byte("Hello"))
	} else {
		conn.Write([]byte("Invalid credentials"))
	}
}

func main() {
	listener, err := net.Listen("tcp", ":5000")
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server is listening on :5000")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go handleConnection(conn)
	}
}
