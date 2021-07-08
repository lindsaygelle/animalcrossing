package cat

import (
	"github.com/lindsaygelle/animalcrossing/animals"
	"github.com/lindsaygelle/animalcrossing/translations"
)

var (
	Cat animals.Animal = cat{}
)

type cat struct{}

func (b cat) Name() string {
	return "Cat"
}

func (b cat) Special() bool {
	return false
}

func (b cat) Translation() translations.Translation {
	return translation{}
}
