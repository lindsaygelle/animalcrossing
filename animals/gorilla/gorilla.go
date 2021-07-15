package gorilla

import (
	"github.com/lindsaygelle/animalcrossing/animals/gorilla/translations"
	"github.com/lindsaygelle/animalcrossing/languages"
)

type Gorilla struct{}

func (g Gorilla) Id() string {
	return "gorilla"
}

func (g Gorilla) Special() bool {
	return false
}

func (g Gorilla) Translations() []languages.Translation {
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
