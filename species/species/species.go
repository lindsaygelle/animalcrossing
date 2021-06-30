package species

type Species interface {
	Name() string
}

type species struct {
	name string
}

func (s species) Name() string {
	return s.name
}

var (
	_ Species = species{}
)
