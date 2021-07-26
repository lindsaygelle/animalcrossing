
package tiger

import (
	a "github.com/lindsaygelle/animalcrossing/animals/tiger"
	s "github.com/lindsaygelle/animalcrossing/species/tiger"
)

// Tiger is an Animal Crossing villager type.
type Tiger struct{}

func (v Tiger) Animal() string {
	return (a.Tiger{}).Id()
}

func (v Tiger) Species() string {
	return (s.Tiger{}).Id()
}
