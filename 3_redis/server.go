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

	redis := Redis{values: make(map[string]string)}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting conn:", err)
			continue
		}

		go handleConn(conn, &redis)
	}
}

func handleConn(conn net.Conn, redis *Redis) {
	defer conn.Close()

	reader := bufio.NewReader(conn)

	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading message:", err)
			return
		}

		conn.Write(handleMessage(message, redis))
	}
}

func handleMessage(message string, redis *Redis) []byte {
	message = strings.TrimSpace(message)
	words := strings.Split(message, " ")

	// fmt.Printf("[SERVER] Received message: ")
	// for i, word := range words {
	// 	fmt.Printf("%d: %s ", i, word)
	// }
	// fmt.Println()

	switch words[0] {
	case "SET":
		return handleSet(words, redis)
	case "GET":
		return handleGet(words, redis)
	case "DELETE":
		return handleDelete(words, redis)
	default:
		fmt.Println("Unknown command.")
	}

	return []byte("Server Received Your Message\n")
}

func handleSet(words []string, redis *Redis) []byte {
	if len(words) < 3 {
		return []byte("Incorrect amount of arguments for a SET command.\n")
	}

	err := redis.Set(words[1], words[2])

	if err != nil {
		return []byte(err.Error() + "\n")
	}

	fmt.Printf("Client set: %s to: %s\n", words[1], words[2])
	return []byte("Set value at: " + words[1] + "\n")
}

func handleGet(words []string, redis *Redis) []byte {
	if len(words) < 2 {
		return []byte("Incorrect amount of arguments for a GET command.\n")
	}

	value, err := redis.Get(words[1])
	if err != nil {
		return []byte(err.Error() + "\n")
	}

	fmt.Printf("Client got: %s at: %s\n", value, words[1])
	return []byte("Value: " + value + "\n")
}

func handleDelete(words []string, redis *Redis) []byte {
	if len(words) < 2 {
		return []byte("Incorrect amount of arguments for a DELETE command.\n")
	}

	err := redis.Delete(words[1])
	if err != nil {
		return []byte(err.Error() + "\n")
	}

	fmt.Printf("Client deleted value at: %s\n", words[1])
	return []byte("Deleted value at: " + words[1] + "\n")
}
