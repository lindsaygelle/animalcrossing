package villager

import (
	"github.com/lindsaygelle/animalcrossing/animal"
	"github.com/lindsaygelle/animalcrossing/astrology"
	"github.com/lindsaygelle/animalcrossing/birthday"
	"github.com/lindsaygelle/animalcrossing/personality"
	"github.com/lindsaygelle/animalcrossing/translation"
)

// Villager is an Animal Crossing villager.
type Villager interface {
	Animal() animal.Animal
	Astrology() (astrology.Astrology, bool)
	Birthday() birthday.Birthday
	Catchphrases() translation.Translations
	Id() string
	Names() translation.Translations
	Personality() (personality.Personality, bool)
	Special() bool
}

// villager implements Villager.
type villager struct {
	animal       animal.Animal
	astrology    astrology.Astrology
	birthday     birthday.Birthday
	catchphrases translation.Translations
	id           string
	names        translation.Translations
	personality  personality.Personality
	special      bool
}

func (v villager) Animal() animal.Animal {
	return v.animal
}

func (v villager) Astrology() (astrology.Astrology, bool) {
	return v.astrology, (v.astrology != nil)
}

func (v villager) Birthday() birthday.Birthday {
	return v.birthday
}

func (v villager) Catchphrases() translation.Translations {
	return v.catchphrases
}

func (v villager) Id() string {
	return v.id
}

func (v villager) Names() translation.Translations {
	return v.names
}

func (v villager) Personality() (personality.Personality, bool) {
	return v.personality, (v.personality != nil)
}

func (v villager) Special() bool {
	return v.special
}

var (
	_ Villager = villager{}
)
