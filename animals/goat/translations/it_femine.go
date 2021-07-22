package translations

import (
	"github.com/lindsaygelle/animalcrossing/genders"
	"github.com/lindsaygelle/animalcrossing/languages/it"
)

type ItFemine struct {
	it.It
	genders.Female
}

func (it ItFemine) Value() string {
	return "Capra"
}