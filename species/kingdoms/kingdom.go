package kingdoms

type Kingdom interface {
	Name() string
}

type kingdom struct {
	name string
}

func (k kingdom) Name() string {
	return k.name
}

var (
	_ Kingdom = kingdom{}
)
