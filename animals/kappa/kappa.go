package kappa

import (
	"github.com/lindsaygelle/animalcrossing/animals/kappa/translations"
	"github.com/lindsaygelle/animalcrossing/languages"
)

type Kappa struct{}

func (k Kappa) Id() string {
	return "kappa"
}

func (k Kappa) Special() bool {
	return true
}

func (k Kappa) Translations() []languages.Translation {
	return []languages.Translation{
		translations.En{}}
}
