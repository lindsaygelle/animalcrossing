package translations

import (
	"github.com/lindsaygelle/animalcrossing/languages/it"
)

type ItFemine struct {
	it.Femine
}

func (i ItFemine) Value() string {
	return "Gatta"
}
