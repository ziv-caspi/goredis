package processors

import (
	"goredis/database"
	"goredis/protocol"
)

type GetProcessor struct{}

func (processor *GetProcessor) Process(databse *database.Database, command protocol.Command) protocol.Result {
	if command.CommandType != protocol.GET || command.GetParams.Key == "" {
		return protocol.Result{Success: false, CommandType: protocol.GET}
	}

	val, ok := databse.Get(command.GetParams.Key)
	if !ok {
		return protocol.Result{Success: false, CommandType: protocol.GET}
	}

	return protocol.Result{Success: true, GetResponse: protocol.GetResponse{Value: val}, CommandType: protocol.GET}
}
