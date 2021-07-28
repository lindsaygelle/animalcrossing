package gayle

import (
	"fmt"
	"time"

	"github.com/lindsaygelle/animalcrossing/astrology/taurus"
	"github.com/lindsaygelle/animalcrossing/villagers"
	"github.com/lindsaygelle/animalcrossing/villagers/alligator"
)

// Gayle is an Animal Crossing villager.
type Gayle struct {
	alligator.Alligator
	taurus.Taurus
}

func (v Gayle) Appearances() {
	// TBC
}

func (v Gayle) Astrology() string {
	return v.Taurus.Id()
}

func (v Gayle) AstrologyIcon() string {
	return v.Taurus.Icon()
}

func (v Gayle) Birthday() uint8 {
	return 17
}

func (v Gayle) BirthdayMonth() time.Month {
	return time.May
}

func (v Gayle) Code() string {
	return fmt.Sprintf("%s%d", v.Alligator.Code(), v.Number())
}

func (v Gayle) Debut() {
	// TBC
}

func (v Gayle) Gender() string {
	return "female"
}

func (v Gayle) Icon() {
	// TBC
}

func (v Gayle) Id() string {
	return "gayle"
}

func (v Gayle) Languages() []interface{ Value() string } {
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

func (v Gayle) Number() uint8 {
	return 6
}

func (v Gayle) Personality() string {
	return "normal"
}

func (v Gayle) Special() bool {
	return false
}

var (
	_ villagers.Villager = (Gayle{})
)
