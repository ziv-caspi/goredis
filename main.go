package main

import (
	"encoding/json"
	"fmt"
	"goredis/backend"
	"goredis/protocol"
	"net"
)

func main() {
	backend := backend.NewBackend()
	seedData(&backend)

	PORT := "0.0.0.0:8080"
	l, err := net.Listen("tcp4", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	fmt.Println("Listening on:", PORT)
	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("new connection from:", c.LocalAddr())
		go handleConnection(c, &backend)
	}
}

func handleConnection(c net.Conn, backend *backend.Backend) {
	defer c.Close()
	decoder := json.NewDecoder(c)
	encoder := json.NewEncoder(c)
	for {
		var command protocol.Command
		err := decoder.Decode(&command)
		if err != nil {
			fmt.Println("eror reading message: ", err)
			result := protocol.Result{Success: false}
			encoder.Encode(&result)
			break
		}

		result := <-backend.RegisterCommand(command)
		encoder.Encode(&result)
	}
}

func seedData(backend *backend.Backend) {
	if backend == nil {
		return
	}

	backend.Database.Set("key1", "abcd")
	backend.Database.Set("key2", "EFG")
	backend.Database.Set("key3", "ABC")
}
