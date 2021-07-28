package cancer

import (
	a "github.com/lindsaygelle/animalcrossing/astrology"
)

// Cancer is an Animal Crossing astrological star sign type.
type Cancer struct{}

func (v Cancer) Icon() string {
	return "♋"
}

func (v Cancer) Id() string {
	return "cancer"
}

var (
	_ a.Astrology = (Cancer{})
)
