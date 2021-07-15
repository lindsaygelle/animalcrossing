package dodo

import (
	"github.com/lindsaygelle/animalcrossing/animals/dodo/translations"
	"github.com/lindsaygelle/animalcrossing/languages"
)

type Dodo struct{}

func (d Dodo) Id() string {
	return "dodo"
}

func (d Dodo) Special() bool {
	return true
}

func (d Dodo) Translations() []languages.Translation {
	return []languages.Translation{
		translations.En{}}
}
