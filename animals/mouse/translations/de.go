package translations

import (
	"github.com/lindsaygelle/animalcrossing/genders"
	"github.com/lindsaygelle/animalcrossing/languages/de"
)

type De struct {
	de.De
	genders.Male
}

func (d De) Value() string {
	return "Maus"
}
