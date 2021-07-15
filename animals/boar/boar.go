package boar

import (
	"github.com/lindsaygelle/animalcrossing/animals/boar/translations"
	"github.com/lindsaygelle/animalcrossing/languages"
)

type Boar struct{}

func (b Boar) Id() string {
	return "boar"
}

func (b Boar) Special() bool {
	return true
}

func (b Boar) Translations() []languages.Translation {
	return []languages.Translation{
		translations.En{}}
}
