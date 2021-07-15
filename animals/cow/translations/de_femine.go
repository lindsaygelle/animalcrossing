package translations

import (
	"github.com/lindsaygelle/animalcrossing/genders"
	"github.com/lindsaygelle/animalcrossing/languages/de"
)

type DeFemine struct {
	de.De
	genders.Female
}

func (d DeFemine) Value() string {
	return "Kuh"
}
