package cow

import (
	"github.com/lindsaygelle/animalcrossing/animals/cow/translations"
	"github.com/lindsaygelle/animalcrossing/languages"
)

type Cow struct{}

func (c Cow) Id() string {
	return "cow"
}

func (c Cow) Special() bool {
	return false
}

func (c Cow) Translations() []languages.Translation {
	return []languages.Translation{
		translations.DeFemine{},
		translations.En{},
		translations.EsFemine{},
		translations.FrFemine{},
		translations.ItFemine{},
		translations.Jp{},
		translations.Ko{},
		translations.Nl{},
		translations.Ru{},
		translations.Zh{}}
}
