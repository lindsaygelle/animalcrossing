package boots

import (
	"time"

	"github.com/lindsaygelle/animalcrossing/villagers"
	"github.com/lindsaygelle/animalcrossing/villagers/alligator"
)

// Boots is an Animal Crossing villager.
type Boots struct {
	alligator.Alligator
}

func (v Boots) Appearances() {
	// TBC
}

func (v Boots) Astrology() string {
	return "leo"
}

func (v Boots) AstrologyIcon() string {
	return "♌"
}

func (v Boots) Birthday() uint8 {
	return 7
}

func (v Boots) BirthdayMonth() time.Month {
	return time.August
}

func (v Boots) Code() string {
	return "crd03"
}

func (v Boots) Debut() {
	// TBC
}

func (v Boots) Gender() string {
	return "male"
}

func (v Boots) Icon() {
	// TBC
}

func (v Boots) Id() string {
	return "boots"
}

func (v Boots) Languages() []interface{ Value() string } {
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

func (v Boots) Number() uint8 {
	return 3
}

func (v Boots) Personality() string {
	return "jock"
}

func (v Boots) Special() bool {
	return false
}

var (
	_ villagers.Villager = (Boots{})
)
