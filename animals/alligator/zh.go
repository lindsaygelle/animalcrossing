package alligator

import z "github.com/lindsaygelle/animalcrossing/languages/zh"

type zh struct {
	z.Zh
}

func (z zh) Value() string { return "鳄鱼" }
