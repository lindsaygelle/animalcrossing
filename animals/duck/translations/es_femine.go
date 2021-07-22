package translations

import (
	"github.com/lindsaygelle/animalcrossing/languages/es"
)

type EsFemine struct {
	es.Femine
}

func (e EsFemine) Id() string {
	return "femine"
}

func (e EsFemine) Value() string {
	return "Pata"
}
