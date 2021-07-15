package frill_necked_lizard

import (
	"github.com/lindsaygelle/animalcrossing/animals/frill_necked_lizard/translations"
	"github.com/lindsaygelle/animalcrossing/languages"
)

type FrillNeckedLizard struct{}

func (f FrillNeckedLizard) Id() string {
	return "frill_necked_lizard"
}

func (f FrillNeckedLizard) Special() bool {
	return true
}

func (f FrillNeckedLizard) Translations() []languages.Translation {
	return []languages.Translation{
		translations.En{}}
}
