package families

type Family interface {
	Name() string
}

type family struct {
	name string
}

func (f family) Name() string {
	return f.name
}

var (
	_ Family = family{}
)
