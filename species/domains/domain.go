package domains

type Domain interface {
	Name() string
}

type domain struct {
	name string
}

func (d domain) Name() string {
	return d.name
}

var (
	_ Domain = domain{}
)
