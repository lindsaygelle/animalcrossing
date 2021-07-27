package del

import (
	"fmt"
	"time"

	"github.com/lindsaygelle/animalcrossing/villagers"
	"github.com/lindsaygelle/animalcrossing/villagers/alligator"
)

// Del is an Animal Crossing villager.
type Del struct {
	alligator.Alligator
}

func (v Del) Appearances() {
	// TBC
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
	return fmt.Sprintf("%s%d", v.Alligator.Code(), v.Number())
}

func (v Del) Debut() {
	// TBC
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

var (
	_ villagers.Villager = (Del{})
)
