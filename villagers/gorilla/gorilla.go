
package gorilla

import (
	a "github.com/lindsaygelle/animalcrossing/animals/gorilla"
	s "github.com/lindsaygelle/animalcrossing/species/gorilla"
)

// Gorilla is an Animal Crossing villager type.
type Gorilla struct{}

func (v Gorilla) Animal() string {
	return (a.Gorilla{}).Id()
}

func (v Gorilla) Species() string {
	return (s.Gorilla{}).Id()
}
