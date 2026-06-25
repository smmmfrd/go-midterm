package redis

import (
	"bufio"
	"fmt"
	"net"
)

// Runs some checks on the Redis
func Run() {
	fmt.Println("hello from the test script")

	conn, err := net.Dial("tcp", "localhost:8090")
	if err != nil {
		fmt.Println("Error connecting to tcp server: ", err)
		return
	}

	defer conn.Close()

	fmt.Fprintf(conn, "hello from run\n")

	reader := bufio.NewReader(conn)
	res, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading response: ", err)
		return
	}

	fmt.Println("Response: ", res)
}

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
			fmt.Println("Error accepting conn: ", err)
			continue
		}

		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	message, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading message: ", err)
		return
	}

	fmt.Printf("Received message: %s\n", message)

	conn.Write([]byte("Message Received\n"))
}
