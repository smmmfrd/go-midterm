package redis

import (
	"bufio"
	"fmt"
	"net"
)

func Client() {
	fmt.Println("")

	conn, err := net.Dial("tcp", "localhost:8090")
	if err != nil {
		fmt.Println("Error connecting to tcp server:", err)
		return
	}

	defer conn.Close()

	fmt.Fprintf(conn, "Hello message from Client\n")

	reader := bufio.NewReader(conn)
	res, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	fmt.Println("[CLIENT] Received response:", res)
}
