package villager

import (
	"github.com/lindsaygelle/animalcrossing/animal"
	"golang.org/x/text/language"
)

type Villager struct {
	animal.Animal

	Code    string
	Id      string
	Name    map[language.Tag]string
	Phrase  map[language.Tag]string
	Special bool
}
