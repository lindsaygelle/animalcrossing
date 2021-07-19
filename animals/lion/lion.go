package lion

import (
	"github.com/lindsaygelle/animalcrossing/animals/lion/translations"
	"github.com/lindsaygelle/animalcrossing/languages"
)

type Lion struct{}

func (l Lion) Id() string {
	return "lion"
}

func (l Lion) Special() bool {
	return false
}

func (l Lion) Translations() []languages.Translation {
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
