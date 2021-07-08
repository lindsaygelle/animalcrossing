package bear_cub

import "github.com/lindsaygelle/animalcrossing/translations"

var (
	_ translations.Language = (chineseTraditional{})
)

type chineseTraditional struct {
	translations.ChineseTraditional
}

func (ct chineseTraditional) Value() string {
	return "鳥"
}
