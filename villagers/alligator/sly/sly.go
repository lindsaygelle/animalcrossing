package sly

import (
	"fmt"
	"time"

	"github.com/lindsaygelle/animalcrossing/astrology/scorpio"
	"github.com/lindsaygelle/animalcrossing/villagers"
	"github.com/lindsaygelle/animalcrossing/villagers/alligator"
)

// Sly is an Animal Crossing villager.
type Sly struct {
	alligator.Alligator
	scorpio.Scorpio
}

func (v Sly) Appearances() {
	// TBC
}

func (v Sly) Astrology() string {
	return v.Scorpio.Id()
}

func (v Sly) AstrologyIcon() string {
	return v.Scorpio.Icon()
}

func (v Sly) Birthday() uint8 {
	return 15
}

func (v Sly) BirthdayMonth() time.Month {
	return time.November
}

func (v Sly) Code() string {
	return fmt.Sprintf("%s%d", v.Alligator.Code(), v.Number())
}

func (v Sly) Debut() {
	// TBC
}

func (v Sly) Gender() string {
	return "male"
}

func (v Sly) Icon() {
	// TBC
}

func (v Sly) Id() string {
	return "sly"
}

func (v Sly) Languages() []interface{ Value() string } {
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

func (v Sly) Number() uint8 {
	return 9
}

func (v Sly) Personality() string {
	return "jock"
}

func (v Sly) Special() bool {
	return false
}

var (
	_ villagers.Villager = (Sly{})
)
