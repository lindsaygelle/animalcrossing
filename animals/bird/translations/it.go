package translations

import (
	"github.com/lindsaygelle/animalcrossing/languages/it"
)

type It struct {
	it.Masculine
}

func (i It) Value() string {
	return "Uccello"
}
