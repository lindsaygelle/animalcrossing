package mouse

import (
	"github.com/lindsaygelle/animalcrossing/animals/mouse/translations"
	"github.com/lindsaygelle/animalcrossing/languages"
)

type Mouse struct{}

func (m Mouse) Id() string {
	return "mouse"
}

func (m Mouse) Special() bool {
	return false
}

func (m Mouse) Translations() []languages.Translation {
	return []languages.Translation{
		translations.De{},
		translations.En{},
		translations.EsFemine{},
		translations.Es{},
		translations.Fr{},
		translations.It{},
		translations.Jp{},
		translations.Ko{},
		translations.Nl{},
		translations.Ru{},
		translations.Zh{}}
}
