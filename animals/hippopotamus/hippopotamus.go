package hippopotamus

import (
	"github.com/lindsaygelle/animalcrossing/animals/hippopotamus/translations"
	"github.com/lindsaygelle/animalcrossing/languages"
)

type Hippopotamus struct{}

func (h Hippopotamus) Id() string {
	return "hippopotamus"
}

func (h Hippopotamus) Special() bool {
	return false
}

func (h Hippopotamus) Translations() []languages.Translation {
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
