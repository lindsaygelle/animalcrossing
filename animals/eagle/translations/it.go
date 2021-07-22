package translations

import (
	"github.com/lindsaygelle/animalcrossing/languages/it"
)

type It struct {
	it.Masculine
}

func (i It) Id() string {
	return "masculine"
}

func (i It) Value() string {
	return "Aquila"
}
