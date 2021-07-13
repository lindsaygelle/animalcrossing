package alligator

import d "github.com/lindsaygelle/animalcrossing/languages/de"

type de struct {
	d.De
}

func (d de) Value() string { return "Krokodil" }
