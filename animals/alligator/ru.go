package alligator

import r "github.com/lindsaygelle/animalcrossing/languages/ru"

type ru struct {
	r.Ru
}

func (r ru) Value() string { return "Крокодил" }
