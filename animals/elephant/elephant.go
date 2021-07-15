package elephant

import (
	"github.com/lindsaygelle/animalcrossing/animals/elephant/translations"
	"github.com/lindsaygelle/animalcrossing/languages"
)

type Elephant struct{}

func (e Elephant) Id() string {
	return "elephant"
}

func (e Elephant) Special() bool {
	return false
}

func (e Elephant) Translations() []languages.Translation {
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
