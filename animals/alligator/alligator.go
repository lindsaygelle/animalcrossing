package alligator

import (
	"github.com/lindsaygelle/animalcrossing/animals/alligator/translations"
	"github.com/lindsaygelle/animalcrossing/languages"
)

type Alligator struct{}

func (a Alligator) Id() string {
	return "alligator"
}

func (a Alligator) Special() bool {
	return false
}

func (a Alligator) Translations() []languages.Translation {
	return []languages.Translation{
		translations.De{},
		translations.En{},
		translations.Es{},
		translations.Fr{},
		translations.It{},
		translations.Jp{},
		translations.Ko{},
		translations.Nl{},
		translations.Ru{},
		translations.Zh{}}
}
