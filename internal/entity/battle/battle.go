package battle

type Battle struct {
	id string
}

func New(id string) *Battle {
	return &Battle{id: id}
}

func (b *Battle) Id() string {
	return b.id
}
