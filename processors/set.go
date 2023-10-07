package processors

import (
	"goredis/database"
	"goredis/protocol"
	"time"
)

type SetProcessor struct{}

func (processor *SetProcessor) Process(databse *database.Database, command protocol.Command) protocol.Result {
	if command.CommandType != protocol.SET || command.SetParams.Key == "" || command.SetParams.Value == "" {
		return protocol.Result{Success: false, CommandType: protocol.SET}
	}

	databse.Set(command.SetParams.Key, command.SetParams.Value, time.Duration(command.SetParams.TtlSeconds))

	return protocol.Result{Success: true, SetReponse: protocol.SetReponse{NewValue: command.SetParams.Value}, CommandType: protocol.SET}
}
