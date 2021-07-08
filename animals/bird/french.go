package bear_cub

import "github.com/lindsaygelle/animalcrossing/translations"

var (
	_ translations.Language = (french{})
)

type french struct {
	translations.French
}

func (f french) Value() string {
	return "Oiseau"
}
