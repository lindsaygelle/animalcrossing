package hobbies

type Hobby interface {
	Name() string
}

type hobby struct {
	name string
}

func (h hobby) Name() string {
	return h.name
}

var (
	_ Hobby = hobby{}
)
