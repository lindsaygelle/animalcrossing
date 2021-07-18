package translations

import (
	"github.com/lindsaygelle/animalcrossing/genders"
	"github.com/lindsaygelle/animalcrossing/languages/fr"
)

type Fr struct {
	fr.Fr
	genders.Male
}

func (f Fr) Value() string {
	return "Hamster"
}
