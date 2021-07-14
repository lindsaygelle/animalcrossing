package translations

import "github.com/lindsaygelle/animalcrossing/languages/de"

type De struct {
	de.De
}

func (d De) Value() string {
	return "Jungbär"
}
