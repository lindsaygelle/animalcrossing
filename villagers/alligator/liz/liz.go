package liz

import (
	"fmt"
	"time"

	"github.com/lindsaygelle/animalcrossing/villagers"
	"github.com/lindsaygelle/animalcrossing/villagers/alligator"
)

// Liz is an Animal Crossing villager.
type Liz struct {
	alligator.Alligator
}

func (v Liz) Appearances() {
	// TBC
}

func (v Liz) Astrology() string {
	return "virgo"
}

func (v Liz) AstrologyIcon() string {
	return "♍"
}

func (v Liz) Birthday() uint8 {
	return 7
}

func (v Liz) BirthdayMonth() time.Month {
	return time.September
}

func (v Liz) Code() string {
	return fmt.Sprintf("%s%d", v.Alligator.Code(), v.Number())
}

func (v Liz) Debut() {
	// TBC
}

func (v Liz) Gender() string {
	return "female"
}

func (v Liz) Icon() {
	// TBC
}

func (v Liz) Id() string {
	return "liz"
}

func (v Liz) Languages() []interface{ Value() string } {
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

func (v Liz) Number() uint8 {
	return 7
}

func (v Liz) Personality() string {
	return "normal"
}

func (v Liz) Special() bool {
	return false
}

var (
	_ villagers.Villager = (Liz{})
)
