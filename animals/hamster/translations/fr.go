package translations

import (
	"github.com/lindsaygelle/animalcrossing/languages/fr"
)

type Fr struct {
	fr.Masculine
}

func (f Fr) Id() string {
	return "masculine"
}

func (f Fr) Value() string {
	return "Hamster"
}
