package dog

import (
	"github.com/lindsaygelle/animalcrossing/animals/dog/translations"
	"github.com/lindsaygelle/animalcrossing/languages"
)

type Dog struct{}

func (d Dog) Id() string {
	return "dog"
}

func (d Dog) Special() bool {
	return false
}

func (d Dog) Translations() []languages.Translation {
	return []languages.Translation{
		translations.De{},
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
		translations.RuFemine{},
		translations.Ru{},
		translations.Zh{}}
}
