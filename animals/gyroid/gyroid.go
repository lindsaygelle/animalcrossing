package gyroid

import (
	"github.com/lindsaygelle/animalcrossing/animals/gyroid/translations"
	"github.com/lindsaygelle/animalcrossing/languages"
)

type Gyroid struct{}

func (g Gyroid) Id() string {
	return "gyroid"
}

func (g Gyroid) Special() bool {
	return true
}

func (g Gyroid) Translations() []languages.Translation {
	return []languages.Translation{
		translations.En{}}
}
