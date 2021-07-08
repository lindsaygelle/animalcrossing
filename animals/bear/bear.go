package bear

import (
	"github.com/lindsaygelle/animalcrossing/animals"
	"github.com/lindsaygelle/animalcrossing/translations"
)

var (
	Bear animals.Animal = bear{}
)

type bear struct{}

func (a bear) Name() string {
	return "Bear"
}

func (a bear) Special() bool {
	return false
}

func (a bear) Translation() translations.Translation {
	return translation{}
}
