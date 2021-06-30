package orders

type Order interface {
	Name() string
}

type order struct {
	name string
}

func (o order) Name() string {
	return o.name
}

var (
	_ Order = order{}
)
