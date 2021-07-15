package translations

import (
	"github.com/lindsaygelle/animalcrossing/genders"
	"github.com/lindsaygelle/animalcrossing/languages/fr"
)

type FrFemine struct {
	fr.Fr
	genders.Female
}

func (f FrFemine) Value() string {
	return "Poule"
}
