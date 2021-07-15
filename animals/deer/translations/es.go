package translations

import (
	"github.com/lindsaygelle/animalcrossing/genders"
	"github.com/lindsaygelle/animalcrossing/languages/es"
)

type Es struct {
	es.Es
	genders.Male
}

func (e Es) Value() string {
	return "Venado"
}
