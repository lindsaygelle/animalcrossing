package horse

import (
	"github.com/lindsaygelle/animalcrossing/animals/horse/translations"
	"github.com/lindsaygelle/animalcrossing/languages"
)

type Horse struct{}

func (h Horse) Id() string {
	return "horse"
}

func (h Horse) Special() bool {
	return false
}

func (h Horse) Translations() []languages.Translation {
	return []languages.Translation{
		translations.DeNeutral{},
		translations.En{},
		translations.EsFemine{},
		translations.Es{},
		translations.FrFemine{},
		translations.Fr{},
		translations.ItFemine{},
		translations.It{},
		translations.Jp{},
		translations.Ko{},
		translations.Nl{},
		translations.Ru{},
		translations.Zh{}}
}
