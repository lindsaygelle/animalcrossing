package octopus

import (
	"github.com/lindsaygelle/animalcrossing/animals/octopus/translations"
	"github.com/lindsaygelle/animalcrossing/languages"
)

type Octopus struct{}

func (o Octopus) Id() string {
	return "octopus"
}

func (o Octopus) Special() bool {
	return false
}

func (o Octopus) Translations() []languages.Translation {
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
