package translations

import (
	"github.com/lindsaygelle/animalcrossing/genders"
	"github.com/lindsaygelle/animalcrossing/languages/de"
)

type DeNeutral struct {
	de.De
	genders.Neutral
}

func (d DeNeutral) Value() string {
	return "Hähnchen"
}
