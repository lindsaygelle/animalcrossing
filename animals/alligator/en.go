package alligator

import e "github.com/lindsaygelle/animalcrossing/languages/en"

type en struct {
	e.En
}

func (e en) Value() string { return "Alligator" }
