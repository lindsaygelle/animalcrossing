package translations

import "github.com/lindsaygelle/animalcrossing/languages/de"

type DeFemine struct {
	de.De
}

func (d DeFemine) Value() string {
	return "Katze"
}
