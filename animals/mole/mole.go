package mole

import (
	"github.com/lindsaygelle/animalcrossing/animals/mole/translations"
	"github.com/lindsaygelle/animalcrossing/languages"
)

type Mole struct{}

func (m Mole) Id() string {
	return "mole"
}

func (m Mole) Special() bool {
	return true
}

func (m Mole) Translations() []languages.Translation {
	return []languages.Translation{
		translations.En{}}
}
