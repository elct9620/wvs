package rpc

type Command struct {
	Name       string      `json:"name"`
	Parameters interface{} `json:"parameters"`
}

func NewCommand(name string, parameters interface{}) *Command {
	return &Command{
		Name:       name,
		Parameters: parameters,
	}
}
