package translations

import (
	"github.com/lindsaygelle/animalcrossing/languages/ru"
)

type RuFemine struct {
	ru.Femine
}

func (r RuFemine) Value() string {
	return "Собака"
}
