package alligator

import (
	"github.com/lindsaygelle/animalcrossing/animals"
	"github.com/lindsaygelle/animalcrossing/translations"
)

var (
	Alligator animals.Animal = alligator{}
)

type alligator struct{}

func (a alligator) Name() string {
	return "Alligator"
}

func (a alligator) Special() bool {
	return false
}

func (a alligator) Translation() translations.Translation {
	return translation{}
}
