package translations

import (
	"github.com/lindsaygelle/animalcrossing/languages/de"
)

type DeNeutral struct {
	de.Neuter
}

func (d DeNeutral) Id() string {
	return "neuter"
}

func (d DeNeutral) Value() string {
	return "Pferd"
}
