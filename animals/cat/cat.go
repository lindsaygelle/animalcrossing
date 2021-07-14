package cat

import (
	"github.com/lindsaygelle/animalcrossing/animals/cat/translations"
	"github.com/lindsaygelle/animalcrossing/languages"
)

type Cat struct{}

func (c Cat) Id() string {
	return "cat"
}

func (c Cat) Special() bool {
	return false
}

func (c Cat) Translations() []languages.Translation {
	return []languages.Translation{
		translations.DeFemine{},
		translations.De{},
		translations.En{},
		translations.EsFemine{},
		translations.Es{},
		translations.FrFemine{},
		translations.ItFemine{},
		translations.It{},
		translations.Jp{},
		translations.Ko{},
		translations.Nl{},
		translations.RuFemine{},
		translations.Ru{},
		translations.Zh{}}
}
