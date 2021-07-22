package translations

import "github.com/lindsaygelle/animalcrossing/languages/de"

type De struct {
	de.Masculine
}

func (d De) Id() string {
	return "masculine"
}

func (d De) Value() string {
	return "Stier"
}
