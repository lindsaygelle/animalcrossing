package aquarius

import (
	a "github.com/lindsaygelle/animalcrossing/astrology"
)

// Aquarius is an Animal Crossing astrological star sign type.
type Aquarius struct{}

func (v Aquarius) Icon() string {
	return "♒"
}

func (v Aquarius) Id() string {
	return "aquarius"
}

var (
	_ a.Astrology = (Aquarius{})
)
