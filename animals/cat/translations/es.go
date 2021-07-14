package translations

import "github.com/lindsaygelle/animalcrossing/languages/es"

type Es struct {
	es.Es
}

func (e Es) Value() string {
	return "Gato"
}
