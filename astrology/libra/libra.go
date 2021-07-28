package libra

import (
	a "github.com/lindsaygelle/animalcrossing/astrology"
)

// Libra is an Animal Crossing astrological star sign type.
type Libra struct{}

func (v Libra) Icon() string {
	return "♎"
}

func (v Libra) Id() string {
	return "libra"
}

var (
	_ a.Astrology = (Libra{})
)
