package alligator

import k "github.com/lindsaygelle/animalcrossing/languages/ko"

type ko struct {
	k.Ko
}

func (k ko) Value() string { return "악어" }
