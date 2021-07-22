package translations

import (
	"github.com/lindsaygelle/animalcrossing/languages/de"
)

type De struct {
	de.Masculine
}

func (d De) Value() string {
	return "Bär"
}
