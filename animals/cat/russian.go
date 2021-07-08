package cat

import "github.com/lindsaygelle/animalcrossing/translations"

var (
	_ translations.Language = (russian{})
)

type russian struct {
	translations.Russian
}

func (r russian) Value() string {
	return "Кот"
}
