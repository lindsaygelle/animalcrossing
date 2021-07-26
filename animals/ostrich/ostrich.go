package ostrich

import (
	"github.com/lindsaygelle/animalcrossing/animals/ostrich/translations"
	"github.com/lindsaygelle/animalcrossing/languages"
)

type Ostrich struct{}

func (o Ostrich) Id() string {
	return "ostrich"
}

func (o Ostrich) Special() bool {
	return false
}

func (o Ostrich) Translations() []languages.Translation {
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
