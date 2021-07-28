package alli

import (
	"fmt"
	"time"

	"github.com/lindsaygelle/animalcrossing/astrology/scorpio"
	"github.com/lindsaygelle/animalcrossing/villagers"
	"github.com/lindsaygelle/animalcrossing/villagers/alligator"
)

// Alli is an Animal Crossing villager.
type Alli struct {
	alligator.Alligator
	scorpio.Scorpio
}

func (v Alli) Appearances() {
	// TBC
}

func (v Alli) Astrology() string {
	return v.Scorpio.Id()
}

func (v Alli) AstrologyIcon() string {
	return v.Scorpio.Icon()
}

func (v Alli) Birthday() uint8 {
	return 8
}

func (v Alli) BirthdayMonth() time.Month {
	return time.November
}

func (v Alli) Code() string {
	return fmt.Sprintf("%s%d", v.Alligator.Code(), v.Number())
}

func (v Alli) Debut() {
	// TBC
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

var (
	_ villagers.Villager = (Alli{})
)
