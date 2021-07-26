
package tortoise

import (
	a "github.com/lindsaygelle/animalcrossing/animals/tortoise"
	s "github.com/lindsaygelle/animalcrossing/species/tortoise"
)

// Tortoise is an Animal Crossing villager type.
type Tortoise struct{}

func (v Tortoise) Animal() string {
	return (a.Tortoise{}).Id()
}

func (v Tortoise) Species() string {
	return (s.Tortoise{}).Id()
}
