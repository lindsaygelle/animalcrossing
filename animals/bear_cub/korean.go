package bear_cub

import "github.com/lindsaygelle/animalcrossing/translations"

var (
	_ translations.Language = (korean{})
)

type korean struct {
	translations.Korean
}

func (k korean) Value() string {
	return "꼬마곰"
}
