package leo

import (
	a "github.com/lindsaygelle/animalcrossing/astrology"
)

// Leo is an Animal Crossing astrological star sign type.
type Leo struct{}

func (v Leo) Icon() string {
	return "♌"
}

func (v Leo) Id() string {
	return "leo"
}

var (
	_ a.Astrology = (Leo{})
)
