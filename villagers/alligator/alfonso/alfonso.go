package alfonso

import (
	"time"

	a "github.com/lindsaygelle/animalcrossing/animals/alligator"
	s "github.com/lindsaygelle/animalcrossing/species/alligator"
)

// Alfonso is an Animal Crossing villager.
//
// Alfonso is an Alligator.
type Alfonso struct{}

func (v Alfonso) Appearances() {
	// TBC
}

func (v Alfonso) Animal() string {
	return (a.Alligator{}).Id()
}

func (v Alfonso) Astrology() string {
	return "gemini"
}

func (v Alfonso) AstrologyIcon() string {
	return "♊"
}

func (v Alfonso) Birthday() uint8 {
	return 9
}

func (v Alfonso) BirthdayMonth() time.Month {
	return time.June
}

func (v Alfonso) Code() string {
	return "crd01"
}

func (v Alfonso) Gender() string {
	return "male"
}

func (v Alfonso) Icon() {
	// TBC
}

func (v Alfonso) Id() string {
	return "alfonso"
}

func (v Alfonso) Languages() {
	// TBC
}

func (v Alfonso) Number() uint8 {
	return 1
}

func (v Alfonso) Personality() string {
	return "lazy"
}

func (v Alfonso) Special() bool {
	return false
}

func (v Alfonso) Species() string {
	return (s.Alligator{}).Id()
}
