package hamster

import (
	"github.com/lindsaygelle/animalcrossing/animals/hamster/translations"
	"github.com/lindsaygelle/animalcrossing/languages"
)

type Hamster struct{}

func (h Hamster) Id() string {
	return "hamster"
}

func (h Hamster) Special() bool {
	return false
}

func (h Hamster) Translations() []languages.Translation {
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
