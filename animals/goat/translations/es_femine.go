package translations

import (
	"github.com/lindsaygelle/animalcrossing/genders"
	"github.com/lindsaygelle/animalcrossing/languages/es"
)

type EsFemine struct {
	es.Es
	genders.Female
}

func (es EsFemine) Value() string {
	return "Cabra"
}
