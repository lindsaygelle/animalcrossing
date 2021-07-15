package deer

import (
	"github.com/lindsaygelle/animalcrossing/animals/deer/translations"
	"github.com/lindsaygelle/animalcrossing/languages"
)

type Deer struct{}

func (d Deer) Id() string {
	return "deer"
}

func (d Deer) Special() bool {
	return false
}

func (d Deer) Translation() []languages.Translation {
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
