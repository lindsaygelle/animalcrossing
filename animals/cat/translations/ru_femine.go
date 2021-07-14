package translations

import "github.com/lindsaygelle/animalcrossing/languages/ru"

type RuFemine struct {
	ru.Ru
}

func (r RuFemine) Value() string {
	return "Кошка"
}
