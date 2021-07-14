package bear_cub

import (
	"github.com/lindsaygelle/animalcrossing/animals/bear_cub/translations"
	"github.com/lindsaygelle/animalcrossing/languages"
)

type BearCub struct{}

func (b BearCub) Id() string {
	return "bear_cub"
}

func (b BearCub) Special() bool {
	return false
}

func (b BearCub) Translations() []languages.Translation {
	return []languages.Translation{
		translations.De{},
		translations.En{},
		translations.Es{},
		translations.EsFemine{},
		translations.Fr{},
		translations.FrFemine{},
		translations.It{},
		translations.ItFemine{},
		translations.Jp{},
		translations.Ko{},
		translations.Nl{},
		translations.Ru{},
		translations.Zh{}}
}
