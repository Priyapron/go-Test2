package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		os.Exit(1)
	}
	defer conn.Close()

	var username, password string
	fmt.Print("Enter username : ")
	fmt.Scan(&username)
	fmt.Print("Enter password : ")
	fmt.Scan(&password)

	data := fmt.Sprintf("%s:%s", username, password)

	_, err = conn.Write([]byte(data))
	if err != nil {
		fmt.Println("Error sending data:", err)
		return
	}

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	serverResponse := string(buffer[:n])
	fmt.Println("Server response:", serverResponse)
}
