package sagittarius

import (
	a "github.com/lindsaygelle/animalcrossing/astrology"
)

// Sagittarius is an Animal Crossing astrological star sign type.
type Sagittarius struct{}

func (v Sagittarius) Icon() string {
	return "♐"
}

func (v Sagittarius) Id() string {
	return "sagittarius"
}

var (
	_ a.Astrology = (Sagittarius{})
)
