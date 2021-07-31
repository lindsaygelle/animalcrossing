package animal

// Animal is an Animal Crossing animal.
type Animal interface {
	Code() (string, bool)
	Id() string
	Fictional() bool
	Names() Names
	Special() bool
	Unicode() (string, bool)
}

type animal struct {
	code      string
	id        string
	fictional bool
	names     names
	special   bool
	unicode   string
}

func (v animal) Code() (string, bool) {
	return v.code, (len(v.code) > 0)
}

func (v animal) Id() string {
	return v.id
}

func (v animal) Fictional() bool {
	return v.fictional
}

func (v animal) Names() Names {
	return v.names
}

func (v animal) Special() bool {
	return v.special
}

func (v animal) Unicode() (string, bool) {
	return v.unicode, (len(v.unicode) > 0)
}

var (
	_ Animal = animal{}
)
