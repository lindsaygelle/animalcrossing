package del

import (
	"time"

	a "github.com/lindsaygelle/animalcrossing/animals/alligator"
	s "github.com/lindsaygelle/animalcrossing/species/alligator"
)

// Del is an Animal Crossing villager.
//
// Del is an Alligator.
type Del struct{}

func (v Del) Appearances() {
	// TBC
}

func (v Del) Animal() string {
	return (a.Alligator{}).Id()
}

func (v Del) Astrology() string {
	return "gemini"
}

func (v Del) AstrologyIcon() string {
	return "♊"
}

func (v Del) Birthday() uint8 {
	return 27
}

func (v Del) BirthdayMonth() time.Month {
	return time.May
}

func (v Del) Code() string {
	return "crd04"
}

func (v Del) Gender() string {
	return "male"
}

func (v Del) Icon() {
	// TBC
}

func (v Del) Id() string {
	return "del"
}

func (v Del) Languages() []interface{ Value() string } {
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

func (v Del) Number() uint8 {
	return 4
}

func (v Del) Personality() string {
	return "cranky"
}

func (v Del) Special() bool {
	return false
}

func (v Del) Species() string {
	return (s.Alligator{}).Id()
}
