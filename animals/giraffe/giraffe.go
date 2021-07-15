package giraffe

import (
	"github.com/lindsaygelle/animalcrossing/animals/giraffe/translations"
	"github.com/lindsaygelle/animalcrossing/languages"
)

type Giraffe struct{}

func (g Giraffe) Id() string {
	return "giraffe"
}

func (g Giraffe) Special() bool {
	return true
}

func (g Giraffe) Translations() []languages.Translation {
	return []languages.Translation{
		translations.En{}}
}
