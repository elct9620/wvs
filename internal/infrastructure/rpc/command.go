package rpc

type Parameter interface {
	isParameter()
}

type Command struct {
	Name       string    `json:"name"`
	Parameters Parameter `json:"parameters"`
}

func NewCommand(name string, parameters Parameter) *Command {
	return &Command{
		Name:       name,
		Parameters: parameters,
	}
}
