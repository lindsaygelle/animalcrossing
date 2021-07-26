
package penguin

import (
	a "github.com/lindsaygelle/animalcrossing/animals/penguin"
	s "github.com/lindsaygelle/animalcrossing/species/penguin"
)

// Penguin is an Animal Crossing villager type.
type Penguin struct{}

func (v Penguin) Animal() string {
	return (a.Penguin{}).Id()
}

func (v Penguin) Species() string {
	return (s.Penguin{}).Id()
}
