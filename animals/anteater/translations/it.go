package translations

import "github.com/lindsaygelle/animalcrossing/languages/it"

type It struct {
	it.It
}

func (i It) Value() string {
	return "Formichiere"
}
