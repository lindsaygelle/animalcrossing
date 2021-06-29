package classes

type Class interface {
	Name() string
}

type class struct {
	name string
}

func (c class) Name() string {
	return c.name
}

var (
	_ Class = class{}
)
