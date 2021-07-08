package bear

import (
	"github.com/lindsaygelle/animalcrossing/animals"
	"github.com/lindsaygelle/animalcrossing/translations"
)

var (
	Bear animals.Animal = bear{}
)

type bear struct{}

func (b bear) Name() string {
	return "Bear"
}

func (b bear) Special() bool {
	return false
}

func (b bear) Translation() translations.Translation {
	return translation{}
}
