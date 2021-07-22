package translations

import (
	"github.com/lindsaygelle/animalcrossing/genders"
	"github.com/lindsaygelle/animalcrossing/languages/de"
)

type DeNeutral struct {
	de.Masculine
	genders.Neutral
}

func (d DeNeutral) Value() string {
	return "Pferd"
}
