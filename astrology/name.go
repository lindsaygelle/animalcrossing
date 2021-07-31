package astrology

import "github.com/lindsaygelle/animalcrossing/translation"

// Name is an Astrological star sign name.
type Name interface {
	translation.Translation
}

// name implements name.
type name struct {
	translation.Translation
}

var (
	_ Name = name{}
)
