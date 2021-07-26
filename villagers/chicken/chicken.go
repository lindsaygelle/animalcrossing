
package chicken

import (
	a "github.com/lindsaygelle/animalcrossing/animals/chicken"
	s "github.com/lindsaygelle/animalcrossing/species/chicken"
)

// Chicken is an Animal Crossing villager type.
type Chicken struct{}

func (v Chicken) Animal() string {
	return (a.Chicken{}).Id()
}

func (v Chicken) Species() string {
	return (s.Chicken{}).Id()
}
