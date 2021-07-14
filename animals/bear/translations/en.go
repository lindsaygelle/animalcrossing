package translations

import (
	"github.com/lindsaygelle/animalcrossing/genders"
	e "github.com/lindsaygelle/animalcrossing/languages/en"
)

type En struct {
	e.En
	genders.None
}

func (e En) Value() string {
	return "Bear"
}
