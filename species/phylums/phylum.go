package phylums

type Phylum interface {
	Name() string
}

type phylum struct {
	name string
}

func (p phylum) Name() string {
	return p.name
}

var (
	_ Phylum = phylum{}
)
