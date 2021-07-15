package chicken

import (
	"github.com/lindsaygelle/animalcrossing/animals/chicken/translations"
	"github.com/lindsaygelle/animalcrossing/languages"
)

type Chicken struct{}

func (c Chicken) Id() string {
	return "chicken"
}

func (c Chicken) Special() bool {
	return false
}

func (c Chicken) Translations() []languages.Translation {
	return []languages.Translation{
		translations.DeNeutral{},
		translations.En{},
		translations.Es{},
		translations.FrFemine{},
		translations.Fr{},
		translations.It{},
		translations.Jp{},
		translations.Ko{},
		translations.Nl{},
		translations.Ru{},
		translations.Zh{}}
}
