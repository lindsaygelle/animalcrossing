package translations

import (
	"github.com/lindsaygelle/animalcrossing/languages/fr"
)

type FrFemine struct {
	fr.Femine
}

func (f FrFemine) Id() string {
	return "femine"
}

func (f FrFemine) Value() string {
	return "Oursonne"
}
