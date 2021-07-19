package translations

import (
	"github.com/lindsaygelle/animalcrossing/genders"
	"github.com/lindsaygelle/animalcrossing/languages/nl"
)

type Nl struct {
	nl.Nl
	genders.None
}

func (n Nl) Value() string {
	return "Paard"
}
