
package seagull

import (
	a "github.com/lindsaygelle/animalcrossing/animals/seagull"
	s "github.com/lindsaygelle/animalcrossing/species/seagull"
)

// Seagull is an Animal Crossing villager type.
type Seagull struct{}

func (v Seagull) Animal() string {
	return (a.Seagull{}).Id()
}

func (v Seagull) Species() string {
	return (s.Seagull{}).Id()
}
