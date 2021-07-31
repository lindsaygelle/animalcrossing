package personality

import "golang.org/x/text/language"

// Names is a collection of personality.Name.
type Names interface {
	ForEach(func(language.Tag, Name))
	Get(language.Tag) (Name, bool)
	Len() int
	Must(language.Tag) Name
}

// names implements Names.
type names map[language.Tag]name

func (v names) ForEach(fn func(language.Tag, Name)) {
	for k, v := range v {
		fn(k, v)
	}
}

func (v names) Get(k language.Tag) (Name, bool) {
	name, ok := v[k]
	return name, ok
}

func (v names) Len() int {
	return len(v)
}

func (v names) Must(k language.Tag) Name {
	name, ok := v.Get(k)
	if !ok {
		panic(k)
	}
	return name
}

var (
	_ Names = names{}
)
