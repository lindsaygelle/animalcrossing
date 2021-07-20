package monkey

import (
	"github.com/lindsaygelle/animalcrossing/animals/monkey/translations"
	"github.com/lindsaygelle/animalcrossing/languages"
)

type Monkey struct{}

func (m Monkey) Id() string    { return "monkey" }
func (m Monkey) Special() bool { return false }
func (m Monkey) Translations() []languages.Translation {
	return []languages.Translation{
		translations.De{},
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
