package species

import "github.com/lindsaygelle/animalcrossing/animals"

type bull struct {
	bovine
}

func (b bull) Animal() string {
	return animals.Bull.Name()
}

var (
	Bull Species = bull{}
)
