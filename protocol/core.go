package protocol

type CommandType uint8

const (
	GET CommandType = iota
	SET
)

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
