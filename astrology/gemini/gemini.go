package gemini

import (
	a "github.com/lindsaygelle/animalcrossing/astrology"
)

// Gemini is an Animal Crossing astrological star sign type.
type Gemini struct{}

func (v Gemini) Icon() string {
	return "♊"
}

func (v Gemini) Id() string {
	return "gemini"
}

var (
	_ a.Astrology = (Gemini{})
)
