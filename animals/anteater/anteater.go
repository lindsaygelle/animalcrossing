package anteater

import (
	"github.com/lindsaygelle/animalcrossing/animals/anteater/translations"
	"github.com/lindsaygelle/animalcrossing/languages"
)

type Anteater struct{}

func (a Anteater) Id() string {
	return "anteater"
}

func (a Anteater) Special() bool {
	return false
}

func (a Anteater) Translations() []languages.Translation {
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
