package redis

import (
	"bufio"
	"fmt"
	"net"
)

var messages = []string{"SET name", "GET name", "DELETE name"}

func Client() {
	fmt.Println("")

	conn, err := net.Dial("tcp", "localhost:8090")
	if err != nil {
		fmt.Println("Error connecting to tcp server:", err)
		return
	}

	defer conn.Close()

	for _, message := range messages {
		sendMessage(message, &conn)
	}
}

func sendMessage(message string, conn *net.Conn) {

	fmt.Fprintf(*conn, "%s\n", message)

	reader := bufio.NewReader(*conn)
	res, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	fmt.Println("[CLIENT] Received response:", res)

}
