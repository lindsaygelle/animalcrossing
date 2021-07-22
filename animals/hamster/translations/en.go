package translations

import (
	"github.com/lindsaygelle/animalcrossing/languages/en"
)

type En struct {
	en.En
}

func (e En) Value() string {
	return "Hamster"
}
