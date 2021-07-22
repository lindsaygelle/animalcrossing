package translations

import (
	"github.com/lindsaygelle/animalcrossing/languages/de"
)

type DeFemine struct {
	de.Femine
}

func (d DeFemine) Value() string {
	return "Kuh"
}
