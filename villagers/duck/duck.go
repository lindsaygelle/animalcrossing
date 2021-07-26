
package duck

import (
	a "github.com/lindsaygelle/animalcrossing/animals/duck"
	s "github.com/lindsaygelle/animalcrossing/species/duck"
)

// Duck is an Animal Crossing villager type.
type Duck struct{}

func (v Duck) Animal() string {
	return (a.Duck{}).Id()
}

func (v Duck) Species() string {
	return (s.Duck{}).Id()
}
