package translations

import (
	"github.com/lindsaygelle/animalcrossing/languages/ru"
)

type Ru struct {
	ru.Masculine
}

func (r Ru) Value() string {
	return "Бегемот"
}
