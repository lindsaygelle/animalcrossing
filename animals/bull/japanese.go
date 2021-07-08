package bull

import "github.com/lindsaygelle/animalcrossing/translations"

var (
	_ translations.Language = (japanese{})
)

type japanese struct {
	translations.Japanese
}

func (j japanese) Value() string {
	return "おすうし"
}
