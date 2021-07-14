package translations

import "github.com/lindsaygelle/animalcrossing/languages/ru"

type Ru struct {
	ru.Ru
}

func (r Ru) Value() string {
	return "Бык"
}
