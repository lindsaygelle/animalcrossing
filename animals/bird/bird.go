package bird

import (
	"github.com/lindsaygelle/animalcrossing/animals/bird/translations"
	"github.com/lindsaygelle/animalcrossing/languages"
)

type Bird struct{}

func (b Bird) Id() string {
	return "bird"
}

func (b Bird) Special() bool {
	return false
}

func (b Bird) Translations() []languages.Translation {
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
