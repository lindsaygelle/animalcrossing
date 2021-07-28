package pisces

import (
	a "github.com/lindsaygelle/animalcrossing/astrology"
)

// Pisces is an Animal Crossing astrological star sign type.
type Pisces struct{}

func (v Pisces) Icon() string {
	return "♓"
}

func (v Pisces) Id() string {
	return "pisces"
}

var (
	_ a.Astrology = (Pisces{})
)
