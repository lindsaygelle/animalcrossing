package cyrus

import (
	"fmt"
	"time"

	"github.com/lindsaygelle/animalcrossing/astrology/aquarius"
	"github.com/lindsaygelle/animalcrossing/villagers"
	"github.com/lindsaygelle/animalcrossing/villagers/alpaca"
)

// Cyrus is an Animal Crossing villager.
type Cyrus struct {
	alpaca.Alpaca
	aquarius.Aquarius
}

func (v Cyrus) Appearances() {
	// TBC
}

func (v Cyrus) Astrology() string {
	return v.Aquarius.Id()
}

func (v Cyrus) AstrologyIcon() string {
	return v.Aquarius.Icon()
}

func (v Cyrus) Birthday() uint8 {
	return 26
}

func (v Cyrus) BirthdayMonth() time.Month {
	return time.January
}

func (v Cyrus) Code() string {
	return fmt.Sprintf("%s%d", v.Alpaca.Code(), v.Number())
}

func (v Cyrus) Debut() {
	// TBC
}

func (v Cyrus) Gender() string {
	return "male"
}

func (v Cyrus) Icon() {
	// TBC
}

func (v Cyrus) Id() string {
	return "cyrus"
}

func (v Cyrus) Languages() []interface{ Value() string } {
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

func (v Cyrus) Number() uint8 {
	return 1
}

func (v Cyrus) Personality() string {
	return "none"
}

func (v Cyrus) Special() bool {
	return true
}

var (
	_ villagers.Villager = (Cyrus{})
)
