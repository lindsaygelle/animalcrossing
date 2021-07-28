package drago

import (
	"fmt"
	"time"

	"github.com/lindsaygelle/animalcrossing/astrology/aquarius"
	"github.com/lindsaygelle/animalcrossing/villagers"
	"github.com/lindsaygelle/animalcrossing/villagers/alligator"
)

// Drago is an Animal Crossing villager.
type Drago struct {
	alligator.Alligator
	aquarius.Aquarius
}

func (v Drago) Appearances() {
	// TBC
}

func (v Drago) Astrology() string {
	return v.Aquarius.Id()
}

func (v Drago) AstrologyIcon() string {
	return v.Aquarius.Icon()
}

func (v Drago) Birthday() uint8 {
	return 12
}

func (v Drago) BirthdayMonth() time.Month {
	return time.February
}

func (v Drago) Code() string {
	return fmt.Sprintf("%s%d", v.Alligator.Code(), v.Number())
}

func (v Drago) Debut() {
	// TBC
}

func (v Drago) Gender() string {
	return "male"
}

func (v Drago) Icon() {
	// TBC
}

func (v Drago) Id() string {
	return "drago"
}

func (v Drago) Languages() []interface{ Value() string } {
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

func (v Drago) Number() uint8 {
	return 5
}

func (v Drago) Personality() string {
	return "lazy"
}

func (v Drago) Special() bool {
	return false
}

var (
	_ villagers.Villager = (Drago{})
)
