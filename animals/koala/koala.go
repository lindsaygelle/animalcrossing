package koala

import (
	"github.com/lindsaygelle/animalcrossing/animals/koala/translations"
	"github.com/lindsaygelle/animalcrossing/languages"
)

type Koala struct{}

func (k Koala) Id() string {
	return "koala"
}

func (k Koala) Special() bool {
	return false
}

func (k Koala) Translations() []languages.Translation {
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
