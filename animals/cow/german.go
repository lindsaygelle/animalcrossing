package cow

import "github.com/lindsaygelle/animalcrossing/translations"

var (
	_ translations.Language = (german{})
)

type german struct {
	translations.German
}

func (g german) Value() string {
	return "Kuh"
}
