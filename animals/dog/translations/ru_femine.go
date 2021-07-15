package translations

import (
	"github.com/lindsaygelle/animalcrossing/genders"
	"github.com/lindsaygelle/animalcrossing/languages/ru"
)

type RuFemine struct {
	ru.Ru
	genders.Female
}

func (r RuFemine) Value() string {
	return "Собака"
}
