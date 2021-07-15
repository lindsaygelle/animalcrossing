package translations

import (
	"github.com/lindsaygelle/animalcrossing/genders"
	"github.com/lindsaygelle/animalcrossing/languages/it"
)

type It struct {
	it.It
	genders.Male
}

func (i It) Value() string {
	return "Rana"
}
