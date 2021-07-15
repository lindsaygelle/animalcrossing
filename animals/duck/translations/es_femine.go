package translations

import (
	"github.com/lindsaygelle/animalcrossing/genders"
	"github.com/lindsaygelle/animalcrossing/languages/es"
)

type EsFemine struct {
	es.Es
	genders.Male
}

func (e EsFemine) Value() string {
	return "Pata"
}
