package chameleon

import (
	"github.com/lindsaygelle/animalcrossing/animals/chameleon/translations"
	"github.com/lindsaygelle/animalcrossing/languages"
)

type Chameleon struct{}

func (c Chameleon) Id() string {
	return "chameleon"
}

func (c Chameleon) Special() bool {
	return true
}

func (c Chameleon) Translations() []languages.Translation {
	return []languages.Translation{
		translations.En{}}
}
