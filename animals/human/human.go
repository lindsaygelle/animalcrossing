package human

import (
	"github.com/lindsaygelle/animalcrossing/animals/human/translations"
	"github.com/lindsaygelle/animalcrossing/languages"
)

type Human struct{}

func (h Human) Id() string {
	return "human"
}

func (h Human) Special() bool {
	return true
}

func (h Human) Translations() []languages.Translation {
	return []languages.Translation{
		translations.En{}}
}
