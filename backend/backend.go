package backend

import (
	"goredis/database"
	"goredis/processors"
	"goredis/protocol"
)

type PendingCommand struct {
	command         protocol.Command
	responseChannel chan protocol.Result
}

type Backend struct {
	Database       *database.Database
	processChannel chan PendingCommand
}

func NewBackend() Backend {
	channel := make(chan PendingCommand, 500)
	db := database.NewDatabase()
	backend := Backend{Database: db, processChannel: channel}
	processQueue := func() {
		for {
			pending := <-backend.processChannel
			println("pretented i handled command named: ", pending.command.CommandType)
			pending.responseChannel <- processors.Process(backend.Database, &pending.command)
		}
	}
	go processQueue()
	return backend
}

func (be *Backend) RegisterCommand(command protocol.Command) chan protocol.Result {
	channel := make(chan protocol.Result)
	pending := PendingCommand{command: command, responseChannel: channel}
	be.processChannel <- pending
	return channel
}
