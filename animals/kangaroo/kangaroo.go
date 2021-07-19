package kangaroo

import (
	"github.com/lindsaygelle/animalcrossing/animals/kangaroo/translations"
	"github.com/lindsaygelle/animalcrossing/languages"
)

type Kangaroo struct{}

func (k Kangaroo) Id() string {
	return "kangaroo"
}

func (k Kangaroo) Special() bool {
	return false
}

func (k Kangaroo) Translations() []languages.Translation {
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
