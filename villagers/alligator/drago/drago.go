package drago

import (
	"time"

	"github.com/lindsaygelle/animalcrossing/villagers"
	"github.com/lindsaygelle/animalcrossing/villagers/alligator"
)

// Drago is an Animal Crossing villager.
type Drago struct {
	alligator.Alligator
}

func (v Drago) Appearances() {
	// TBC
}

func (v Drago) Astrology() string {
	return "aquarius"
}

func (v Drago) AstrologyIcon() string {
	return "♒"
}

func (v Drago) Birthday() uint8 {
	return 12
}

func (v Drago) BirthdayMonth() time.Month {
	return time.February
}

func (v Drago) Code() string {
	return "crd05"
}

func (v Drago) Debut() {
	// TBC
}

func (v Drago) Gender() string {
	return "male"
}

func (v Drago) Icon() {
	// TBC
}

func (v Drago) Id() string {
	return "drago"
}

func (v Drago) Languages() []interface{ Value() string } {
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

func (v Drago) Number() uint8 {
	return 5
}

func (v Drago) Personality() string {
	return "lazy"
}

func (v Drago) Special() bool {
	return false
}

var (
	_ villagers.Villager = (Drago{})
)
