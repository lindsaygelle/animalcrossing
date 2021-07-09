package cat

import (
	"github.com/lindsaygelle/animalcrossing/animals"
	"github.com/lindsaygelle/animalcrossing/translations"
)

var (
	Cat animals.Animal = cat{}
)

type cat struct{}

func (c cat) Name() string {
	return "Cat"
}

func (c cat) Special() bool {
	return false
}

func (c cat) Translation() translations.Translation {
	return translation{}
}
