package aries

import (
	a "github.com/lindsaygelle/animalcrossing/astrology"
)

// Aries is an Animal Crossing astrological star sign type.
type Aries struct{}

func (v Aries) Icon() string {
	return "♈"
}

func (v Aries) Id() string {
	return "aries"
}

var (
	_ a.Astrology = (Aries{})
)
