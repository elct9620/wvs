package rpc

type Command struct {
	Name string `json:"name"`
}

func NewCommand(name string) *Command {
	return &Command{
		Name: name,
	}
}
