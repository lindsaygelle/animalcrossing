
package squirrel

import (
	a "github.com/lindsaygelle/animalcrossing/animals/squirrel"
	s "github.com/lindsaygelle/animalcrossing/species/squirrel"
)

// Squirrel is an Animal Crossing villager type.
type Squirrel struct{}

func (v Squirrel) Animal() string {
	return (a.Squirrel{}).Id()
}

func (v Squirrel) Species() string {
	return (s.Squirrel{}).Id()
}
