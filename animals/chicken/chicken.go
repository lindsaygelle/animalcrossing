package chicken

import (
	"github.com/lindsaygelle/animalcrossing/animals"
	"github.com/lindsaygelle/animalcrossing/translations"
)

var (
	Chicken animals.Animal = chicken{}
)

type chicken struct{}

func (c chicken) Name() string {
	return "Chicken"
}

func (c chicken) Special() bool {
	return false
}

func (c chicken) Translation() translations.Translation {
	return translation{}
}
