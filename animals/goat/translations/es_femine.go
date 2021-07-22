package translations

import (
	"github.com/lindsaygelle/animalcrossing/languages/es"
)

type EsFemine struct {
	es.Femine
}

func (es EsFemine) Value() string {
	return "Cabra"
}
