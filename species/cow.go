package species

import "github.com/lindsaygelle/animalcrossing/animals"

type cow struct {
	bovine
}

func (c cow) Animal() string { return animals.Cow.Name() }

var (
	Cow Species = cow{}
)
