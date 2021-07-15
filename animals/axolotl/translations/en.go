package translations

import (
	"github.com/lindsaygelle/animalcrossing/genders"
	"github.com/lindsaygelle/animalcrossing/languages/en"
)

type En struct {
	en.En
	genders.Male
}

func (e En) Value() string {
	return "Axolotl"
}
