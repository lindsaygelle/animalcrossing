package cow

import "github.com/lindsaygelle/animalcrossing/translations"

var (
	_ translations.Language = (dutch{})
)

type dutch struct {
	translations.Dutch
}

func (d dutch) Value() string {
	return "Koe"
}
