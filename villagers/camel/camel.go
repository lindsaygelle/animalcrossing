
package camel

import (
	a "github.com/lindsaygelle/animalcrossing/animals/camel"
	s "github.com/lindsaygelle/animalcrossing/species/camel"
)

// Camel is an Animal Crossing villager type.
type Camel struct{}

func (v Camel) Animal() string {
	return (a.Camel{}).Id()
}

func (v Camel) Species() string {
	return (s.Camel{}).Id()
}
