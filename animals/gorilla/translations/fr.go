package translations

import (
	"github.com/lindsaygelle/animalcrossing/languages/fr"
)

type Fr struct {
	fr.Masculine
}

func (f Fr) Value() string {
	return "Gorille"
}
