package bear_cub

import (
	"github.com/lindsaygelle/animalcrossing/animals"
	"github.com/lindsaygelle/animalcrossing/translations"
)

var (
	Bird animals.Animal = bird{}
)

type bird struct{}

func (b bird) Name() string {
	return "Bird"
}

func (b bird) Special() bool {
	return false
}

func (b bird) Translation() translations.Translation {
	return translation{}
}
