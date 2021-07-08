package anteater

import "github.com/lindsaygelle/animalcrossing/translations"

var (
	_ translations.Language = (italian{})
)

type italian struct {
	translations.Italian
}

func (i italian) Value() string {
	return "Formichiere"
}
