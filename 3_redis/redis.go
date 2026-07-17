package redis

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"sync"
)

type Redis struct {
	mu     sync.Mutex
	values map[string]string
}

func (r *Redis) Set(key string, value string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.values[key]; !ok {
		r.values[key] = value
	} else {
		return fmt.Errorf("Value already exists.")
	}

	return nil
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

	redis := &Redis{values: map[string]string{}}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting conn:", err)
			continue
		}

		go handleConn(conn, redis)
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
	words := strings.Split(message, " ")

	// fmt.Printf("[SERVER] Received message: ")
	// for i, word := range words {
	// 	fmt.Printf("%d: %s ", i, word)
	// }
	// fmt.Println()

	switch words[0] {
	case "SET":
		return handleSet(words, redis)
	// case "GET":
	// 	return handleGet(words)
	// case "DELETE":
	// 	fmt.Println("Client wants to delete a value")
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

// func handleGet(words []string, redis *Redis) []byte {
// 	if len(words) < 2 {
// 		return []byte("Incorrect amount of arguments for a GET command.\n")
// 	}

// 	if _, ok := redis.Values[words[1]]; ok {
// 		fmt.Printf("Client got: %s at: %s\n", redis[words[1]], words[1])
// 	} else {
// 		return []byte("Value does not exist.\n")
// 	}

// 	return []byte("Value: " + redis[words[1]])
// }
