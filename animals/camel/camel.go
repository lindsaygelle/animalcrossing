package camel

import (
	"github.com/lindsaygelle/animalcrossing/animals/camel/translations"
	"github.com/lindsaygelle/animalcrossing/languages"
)

type Camel struct{}

func (c Camel) Id() string {
	return "camel"
}

func (c Camel) Special() bool {
	return true
}

func (c Camel) Translations() []languages.Translation {
	return []languages.Translation{
		translations.En{}}
}
