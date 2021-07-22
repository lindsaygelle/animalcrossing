package translations

import (
	"github.com/lindsaygelle/animalcrossing/languages/es"
)

type EsFemine struct {
	es.Masculine
}

func (e EsFemine) Value() string {
	return "Pata"
}
