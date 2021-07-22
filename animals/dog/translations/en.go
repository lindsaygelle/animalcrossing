package translations

import (
	"github.com/lindsaygelle/animalcrossing/languages/en"
)

type En struct {
	en.En
}

func (e En) Id() string {
	return "none"
}

func (e En) Value() string {
	return "Dog"
}
