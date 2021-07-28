package pironkon

import (
	"fmt"
	"time"

	"github.com/lindsaygelle/animalcrossing/astrology/cancer"
	"github.com/lindsaygelle/animalcrossing/villagers"
	"github.com/lindsaygelle/animalcrossing/villagers/alligator"
)

// Pironkon is an Animal Crossing villager.
type Pironkon struct {
	alligator.Alligator
	cancer.Cancer
}

func (v Pironkon) Appearances() {
	// TBC
}

func (v Pironkon) Astrology() string {
	return v.Cancer.Id()
}

func (v Pironkon) AstrologyIcon() string {
	return v.Cancer.Icon()
}

func (v Pironkon) Birthday() uint8 {
	return 26
}

func (v Pironkon) BirthdayMonth() time.Month {
	return time.June
}

func (v Pironkon) Code() string {
	return fmt.Sprintf("%s%d", v.Alligator.Code(), v.Number())
}

func (v Pironkon) Debut() {
	// TBC
}

func (v Pironkon) Gender() string {
	return "male"
}

func (v Pironkon) Icon() {
	// TBC
}

func (v Pironkon) Id() string {
	return "pironkon"
}

func (v Pironkon) Languages() []interface{ Value() string } {
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

func (v Pironkon) Number() uint8 {
	return 8
}

func (v Pironkon) Personality() string {
	return "lazy"
}

func (v Pironkon) Special() bool {
	return false
}

var (
	_ villagers.Villager = (Pironkon{})
)
