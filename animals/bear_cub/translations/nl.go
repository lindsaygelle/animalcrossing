package translations

import "github.com/lindsaygelle/animalcrossing/languages/nl"

type Nl struct {
	nl.Nl
}

func (n Nl) Value() string {
	return "Berenjong"
}
