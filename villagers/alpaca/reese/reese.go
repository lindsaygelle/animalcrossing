package reese

import (
	"fmt"
	"time"

	"github.com/lindsaygelle/animalcrossing/astrology/cancer"
	"github.com/lindsaygelle/animalcrossing/villagers"
	"github.com/lindsaygelle/animalcrossing/villagers/alpaca"
)

// Reese is an Animal Crossing villager.
type Reese struct {
	alpaca.Alpaca
	cancer.Cancer
}

func (v Reese) Appearances() {
	// TBC
}

func (v Reese) Astrology() string {
	return v.Cancer.Id()
}

func (v Reese) AstrologyIcon() string {
	return v.Cancer.Icon()
}

func (v Reese) Birthday() uint8 {
	return 5
}

func (v Reese) BirthdayMonth() time.Month {
	return time.July
}

func (v Reese) Code() string {
	return fmt.Sprintf("%s%d", v.Alpaca.Code(), v.Number())
}

func (v Reese) Debut() {
	// TBC
}

func (v Reese) Gender() string {
	return "female"
}

func (v Reese) Icon() {
	// TBC
}

func (v Reese) Id() string {
	return "reese"
}

func (v Reese) Languages() []interface{ Value() string } {
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

func (v Reese) Number() uint8 {
	return 2
}

func (v Reese) Personality() string {
	return "none"
}

func (v Reese) Special() bool {
	return true
}

var (
	_ villagers.Villager = (Reese{})
)
