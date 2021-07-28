package capricorn

import (
	a "github.com/lindsaygelle/animalcrossing/astrology"
)

// Capricorn is an Animal Crossing astrological star sign type.
type Capricorn struct{}

func (v Capricorn) Icon() string {
	return "♑"
}

func (v Capricorn) Id() string {
	return "capricorn"
}

var (
	_ a.Astrology = (Capricorn{})
)
