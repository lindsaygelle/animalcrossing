package goat

import (
	"github.com/lindsaygelle/animalcrossing/animals/goat/translations"
	"github.com/lindsaygelle/animalcrossing/languages"
)

type Goat struct{}

func (g Goat) Id() string {
	return "goat"
}

func (g Goat) Special() bool {
	return false
}

func (g Goat) Translation() []languages.Translation {
	return []languages.Translation{
		translations.De{},
		translations.En{},
		translations.EsFemine{},
		translations.Es{},
		translations.FrFemine{},
		translations.Fr{},
		translations.ItFemine{},
		translations.It{},
		translations.Jp{},
		translations.Ko{},
		translations.Nl{},
		translations.Ru{},
		translations.Zh{}}
}
