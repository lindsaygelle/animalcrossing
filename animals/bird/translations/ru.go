package translations

import (
	"github.com/lindsaygelle/animalcrossing/genders"
	"github.com/lindsaygelle/animalcrossing/languages/ru"
)

type Ru struct {
	ru.Ru
	genders.Male
}

func (r Ru) Value() string {
	return "Птица"
}
