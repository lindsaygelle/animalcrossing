package translations

import "github.com/lindsaygelle/animalcrossing/languages/fr"

type FrFemine struct {
	fr.Fr
}

func (f FrFemine) Value() string {
	return "Oursonne"
}
