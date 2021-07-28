package taurus

import (
	a "github.com/lindsaygelle/animalcrossing/astrology"
)

// Taurus is an Animal Crossing astrological star sign type.
type Taurus struct{}

func (v Taurus) Icon() string {
	return "♉"
}

func (v Taurus) Id() string {
	return "taurus"
}

var (
	_ a.Astrology = (Taurus{})
)
