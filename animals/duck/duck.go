package duck

import (
	"github.com/lindsaygelle/animalcrossing/animals/duck/translations"
	"github.com/lindsaygelle/animalcrossing/languages"
)

type Duck struct{}

func (d Duck) Id() string {
	return "duck"
}

func (d Duck) Special() bool {
	return false
}

func (d Duck) Translations() []languages.Translation {
	return []languages.Translation{
		translations.DeFemine{},
		translations.En{},
		translations.EsFemine{},
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
