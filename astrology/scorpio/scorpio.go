package scorpio

import (
	a "github.com/lindsaygelle/animalcrossing/astrology"
)

// Scorpio is an Animal Crossing astrological star sign type.
type Scorpio struct{}

func (v Scorpio) Icon() string {
	return "♏"
}

func (v Scorpio) Id() string {
	return "scorpio"
}

var (
	_ a.Astrology = (Scorpio{})
)
