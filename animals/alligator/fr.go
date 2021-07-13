package alligator

import f "github.com/lindsaygelle/animalcrossing/languages/fr"

type fr struct {
	f.Fr
}

func (f fr) Value() string { return "Crocodile" }
