package alligator

import j "github.com/lindsaygelle/animalcrossing/languages/jp"

type jp struct {
	j.Jp
}

func (j jp) Value() string { return "ワニ" }
