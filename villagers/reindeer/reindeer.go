
package reindeer

import (
	a "github.com/lindsaygelle/animalcrossing/animals/reindeer"
	s "github.com/lindsaygelle/animalcrossing/species/reindeer"
)

// Reindeer is an Animal Crossing villager type.
type Reindeer struct{}

func (v Reindeer) Animal() string {
	return (a.Reindeer{}).Id()
}

func (v Reindeer) Species() string {
	return (s.Reindeer{}).Id()
}
