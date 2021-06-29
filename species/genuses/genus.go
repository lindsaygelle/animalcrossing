package genuses

type Genus interface {
	Name() string
}

type genus struct {
	name string
}

func (g genus) Name() string {
	return g.name
}

var (
	_ Genus = genus{}
)
