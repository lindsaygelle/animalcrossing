package bear

import (
	"github.com/lindsaygelle/animalcrossing/animals/bear/translations"
	"github.com/lindsaygelle/animalcrossing/languages"
)

type Bear struct{}

func (b Bear) Id() string {
	return "bear"
}

func (b Bear) Special() bool {
	return false
}

func (b Bear) Translations() []languages.Translation {
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
