package villager

import (
	"github.com/lindsaygelle/animalcrossing/animal"
	"golang.org/x/text/language"
)

type Villager struct {
	animal.Animal

	Appearance  Appearance
	Birthday    Birthday
	Code        string
	Gender      string
	Id          string
	Key         Key
	Name        map[language.Tag]string
	Personality string
	Phrase      map[language.Tag]string
	Special     bool
}
