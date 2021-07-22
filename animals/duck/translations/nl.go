package translations

import (
	"github.com/lindsaygelle/animalcrossing/languages/nl"
)

type Nl struct {
	nl.Nl
}

func (n Nl) Id() string {
	return "none"
}

func (n Nl) Value() string {
	return "Eend"
}
