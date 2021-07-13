package alligator

import n "github.com/lindsaygelle/animalcrossing/languages/nl"

type nl struct {
	n.Nl
}

func (n nl) Value() string { return "Krokodil" }
