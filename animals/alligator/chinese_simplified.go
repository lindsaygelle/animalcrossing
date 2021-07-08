package alligator

import "github.com/lindsaygelle/animalcrossing/translations"

var (
	_ translations.Language = (chineseSimplified{})
)

type chineseSimplified struct {
	translations.ChineseSimplified
}

func (cs chineseSimplified) Value() string {
	return "鳄鱼"
}
