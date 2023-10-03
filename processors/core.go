package processors

import (
	"errors"
	"fmt"
	"goredis/database"
	"goredis/protocol"
)

type CommandProcessor interface {
	Process(databse *database.Database, command protocol.Command) protocol.Result
}

func resolveProcessor(commandType protocol.CommandType) (CommandProcessor, error) {
	switch commandType {
	case protocol.GET:
		return &GetProcessor{}, nil
	case protocol.SET:
		return &SetProcessor{}, nil
	}

	return nil, errors.New("no processor for this command type")
}

func Process(database *database.Database, command *protocol.Command) protocol.Result {
	processor, error := resolveProcessor(command.CommandType)
	if error != nil {
		fmt.Println(error)
		return protocol.Result{CommandType: command.CommandType, Success: false}
	}

	return processor.Process(database, *command)
}
