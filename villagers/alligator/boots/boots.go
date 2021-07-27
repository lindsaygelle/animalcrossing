package boots

import (
	"time"

	a "github.com/lindsaygelle/animalcrossing/animals/alligator"
	s "github.com/lindsaygelle/animalcrossing/species/alligator"
)

// Boots is an Animal Crossing villager.
//
// Boots is an Alligator.
type Boots struct{}

func (v Boots) Appearances() {
	// TBC
}

func (v Boots) Animal() string {
	return (a.Alligator{}).Id()
}

func (v Boots) Astrology() string {
	return "leo"
}

func (v Boots) AstrologyIcon() string {
	return "♌"
}

func (v Boots) Birthday() uint8 {
	return 7
}

func (v Boots) BirthdayMonth() time.Month {
	return time.August
}

func (v Boots) Code() string {
	return "crd02"
}

func (v Boots) Gender() string {
	return "male"
}

func (v Boots) Icon() {
	// TBC
}

func (v Boots) Id() string {
	return "boots"
}

func (v Boots) Languages() {
	// TBC
}

func (v Boots) Personality() string {
	return "jock"
}

func (v Boots) Special() bool {
	return false
}

func (v Boots) Species() string {
	return (s.Alligator{}).Id()
}
