package personality

import t "github.com/lindsaygelle/animalcrossing/translation"

// Name is a personality name.
type Name interface {
	t.Translation
}

// name implements name.
type name struct {
	t.Translation
}

var (
	_ Name = name{}
)
