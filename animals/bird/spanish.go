package bear_cub

import "github.com/lindsaygelle/animalcrossing/translations"

var (
	_ translations.Language = (spanish{})
)

type spanish struct {
	translations.Spanish
}

func (s spanish) Value() string {
	return "Pájaro"
}
