package hedgehog

import (
	"github.com/lindsaygelle/animalcrossing/animals/hedgehog/translations"
	"github.com/lindsaygelle/animalcrossing/languages"
)

type Hedgehog struct{}

func (h Hedgehog) Id() string {
	return "hedgehog"
}

func (h Hedgehog) Special() bool {
	return true
}

func (h Hedgehog) Value() []languages.Translation {
	return []languages.Translation{
		translations.En{}}
}
