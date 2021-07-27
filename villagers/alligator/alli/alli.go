package alli

import (
	"time"

	a "github.com/lindsaygelle/animalcrossing/animals/alligator"
	s "github.com/lindsaygelle/animalcrossing/species/alligator"
)

// Alli is an Animal Crossing villager.
//
// Alli is an Alligator.
type Alli struct{}

func (v Alli) Appearances() {
	// TBC
}

func (v Alli) Animal() string {
	return (a.Alligator{}).Id()
}

func (v Alli) Astrology() string {
	return "scorpio"
}

func (v Alli) AstrologyIcon() string {
	return "♏"
}

func (v Alli) Birthday() uint8 {
	return 8
}

func (v Alli) BirthdayMonth() time.Month {
	return time.November
}

func (v Alli) Code() string {
	return "crd02"
}

func (v Alli) Gender() string {
	return "female"
}

func (v Alli) Icon() {
	// TBC
}

func (v Alli) Id() string {
	return "alli"
}

func (v Alli) Languages() []interface{ Value() string } {
	return []interface{ Value() string }{
		chineseSimplified{},
		chineseTraditional{},
		dutch{},
		english{},
		french{},
		frenchQuebec{},
		german{},
		italian{},
		japanese{},
		korean{},
		polish{},
		portuguese{},
		russian{},
		spanish{},
		spanishLatinAmerica{}}
}

func (v Alli) Number() uint8 {
	return 2
}

func (v Alli) Personality() string {
	return "snooty"
}

func (v Alli) Special() bool {
	return false
}

func (v Alli) Species() string {
	return (s.Alligator{}).Id()
}
