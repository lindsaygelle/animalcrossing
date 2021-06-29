package conservations

type Conservation interface {
	Name() string
}

type conservation struct {
	name string
}

func (c conservation) Name() string {
	return c.name
}

var (
	_ Conservation = conservation{}
)
