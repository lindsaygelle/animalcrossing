package translations

import (
	"github.com/lindsaygelle/animalcrossing/languages/it"
)

type ItFemine struct {
	it.Femine
}

func (it ItFemine) Value() string {
	return "Capra"
}
