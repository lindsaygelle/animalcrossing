package frog

import (
	"github.com/lindsaygelle/animalcrossing/animals/frog/translations"
	"github.com/lindsaygelle/animalcrossing/languages"
)

type Frog struct{}

func (f Frog) Id() string {
	return "frog"
}

func (f Frog) Special() bool {
	return false
}

func (f Frog) Translations() []languages.Translation {
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
