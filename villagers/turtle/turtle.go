
package turtle

import (
	a "github.com/lindsaygelle/animalcrossing/animals/turtle"
	s "github.com/lindsaygelle/animalcrossing/species/turtle"
)

// Turtle is an Animal Crossing villager type.
type Turtle struct{}

func (v Turtle) Animal() string {
	return (a.Turtle{}).Id()
}

func (v Turtle) Species() string {
	return (s.Turtle{}).Id()
}
