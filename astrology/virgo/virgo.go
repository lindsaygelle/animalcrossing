package virgo

import (
	a "github.com/lindsaygelle/animalcrossing/astrology"
)

// Virgo is an Animal Crossing astrological star sign type.
type Virgo struct{}

func (v Virgo) Icon() string {
	return "♍"
}

func (v Virgo) Id() string {
	return "virgo"
}

var (
	_ a.Astrology = (Virgo{})
)
