package eagle

import (
	"github.com/lindsaygelle/animalcrossing/animals/eagle/translations"
	"github.com/lindsaygelle/animalcrossing/languages"
)

type Eagle struct{}

func (e Eagle) Id() string {
	return "eagle"
}

func (e Eagle) Special() bool {
	return false
}

func (e Eagle) Translations() []languages.Translation {
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
