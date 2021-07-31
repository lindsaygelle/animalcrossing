package villager

import (
	"reflect"

	"github.com/lindsaygelle/animalcrossing/astrology"
	"github.com/lindsaygelle/animalcrossing/personality"
)

// Villager is an Animal Crossing villager.
type Villager interface {
	Astrology() (astrology.Astrology, bool)
	Catchphrases() Catchphrases
	Id() string
	Names() Names
	Personality() (personality.Personality, bool)
	Special() bool
}

// villager implements Villager.
type villager struct {
	astrology    astrology.Astrology
	catchphrases catchphrases
	id           string
	names        names
	personality  personality.Personality
	special      bool
}

func (v villager) Astrology() (astrology.Astrology, bool) {
	return v.astrology, (reflect.ValueOf(v.astrology).IsZero())
}

func (v villager) Catchphrases() Catchphrases {
	return v.catchphrases
}

func (v villager) Id() string {
	return v.id
}

func (v villager) Names() Names {
	return v.names
}

func (v villager) Personality() (personality.Personality, bool) {
	return v.personality, (reflect.ValueOf(v.personality).IsZero())
}

func (v villager) Special() bool {
	return v.special
}

var (
	_ Villager = villager{}
)
