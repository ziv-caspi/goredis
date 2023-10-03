package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
)

func main() {
	fmt.Println("I am your client")

	servAddr := "localhost:8080"
	tcpAddr, err := net.ResolveTCPAddr("tcp", servAddr)
	if err != nil {
		println("ResolveTCPAddr failed:", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		println("Dial failed:", err.Error())
		os.Exit(1)
	}

	for {
		command := Command{CommandType: GET, GetParams: GetParams{Key: "key1"}}
		encoder := json.NewEncoder(conn)
		decoder := json.NewDecoder(conn)
		encoder.Encode(command)

		if err != nil {
			println("Write to server failed:", err.Error())
			os.Exit(1)
		}

		println("write to server command")
		var response Result
		decoder.Decode(&response)
		if err != nil {
			println("Write to server failed:", err.Error())
			os.Exit(1)
		}

		println("reply from server=", response.Success, response.GetResponse.Value)
	}

	conn.Close()
}

type CommandType uint8

const (
	GET CommandType = iota
	SET
)

type GetParams struct {
	Key string
}
type GetResponse struct {
	Value string
}

type SetParams struct {
	Key   string
	Value string
}
type SetReponse struct {
	NewValue string
}

type Command struct {
	CommandType CommandType
	// only one of the fields below should be initialized
	GetParams GetParams
	SetParams SetParams
}

type Result struct {
	CommandType CommandType
	Success     bool
	// only one of the fields below should be initialized
	GetResponse GetResponse
	SetReponse  SetReponse
}
