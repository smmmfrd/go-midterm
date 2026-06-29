package redis

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

// Starts the Redis
func Start() {
	listener, err := net.Listen("tcp", "localhost:8090")
	if err != nil {
		fmt.Println("Error creating tcp server: ", err)
		return
	}

	defer listener.Close()

	fmt.Println("Started TCP server at localhost:8090")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting conn:", err)
			continue
		}

		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)

	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading message:", err)
			return
		}

		conn.Write(handleMessage(message))

	}
}

func handleMessage(message string) []byte {
	words := strings.Split(message, " ")

	fmt.Printf("[SERVER] Received message: ")
	for i, word := range words {
		fmt.Printf("%d: %s ", i, word)
	}
	fmt.Println()

	return []byte("Server Received Your Message\n")
}
