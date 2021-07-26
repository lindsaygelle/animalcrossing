
package bird

import (
	a "github.com/lindsaygelle/animalcrossing/animals/bird"
	s "github.com/lindsaygelle/animalcrossing/species/bird"
)

// Bird is an Animal Crossing villager type.
type Bird struct{}

func (v Bird) Animal() string {
	return (a.Bird{}).Id()
}

func (v Bird) Species() string {
	return (s.Bird{}).Id()
}
