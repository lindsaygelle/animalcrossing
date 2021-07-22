package translations

import (
	"github.com/lindsaygelle/animalcrossing/languages/es"
)

type Es struct {
	es.Masculine
}

func (e Es) Value() string {
	return "Osezno"
}
