package fox

import (
	"github.com/lindsaygelle/animalcrossing/animals/fox/translations"
	"github.com/lindsaygelle/animalcrossing/languages"
)

type Fox struct{}

func (f Fox) Id() string {
	return "fox"
}

func (f Fox) Special() bool {
	return true
}

func (f Fox) Translations() []languages.Translation {
	return []languages.Translation{
		translations.En{}}
}
