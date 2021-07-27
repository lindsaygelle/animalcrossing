package alfonso

import (
	"time"

	"github.com/lindsaygelle/animalcrossing/villagers"
	"github.com/lindsaygelle/animalcrossing/villagers/alligator"
)

// Alfonso is an Animal Crossing villager.
type Alfonso struct {
	alligator.Alligator
}

func (v Alfonso) Appearances() {
	// TBC
}

func (v Alfonso) Astrology() string {
	return "gemini"
}

func (v Alfonso) AstrologyIcon() string {
	return "♊"
}

func (v Alfonso) Birthday() uint8 {
	return 9
}

func (v Alfonso) BirthdayMonth() time.Month {
	return time.June
}

func (v Alfonso) Code() string {
	return "crd01"
}

func (v Alfonso) Debut() {
	// TBC
}

func (v Alfonso) Gender() string {
	return "male"
}

func (v Alfonso) Icon() {
	// TBC
}

func (v Alfonso) Id() string {
	return "alfonso"
}

func (v Alfonso) Languages() []interface{ Value() string } {
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

func (v Alfonso) Number() uint8 {
	return 1
}

func (v Alfonso) Personality() string {
	return "lazy"
}

func (v Alfonso) Special() bool {
	return false
}

var (
	_ villagers.Villager = (Alfonso{})
)
