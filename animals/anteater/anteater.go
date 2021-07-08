package anteater

import (
	"github.com/lindsaygelle/animalcrossing/animals"
	"github.com/lindsaygelle/animalcrossing/translations"
)

var (
	Anteater animals.Animal = anteater{}
)

type anteater struct{}

func (a anteater) Name() string {
	return "Anteater"
}

func (a anteater) Special() bool {
	return false
}

func (a anteater) Translation() translations.Translation {
	return translation{}
}
