
package mouse

import (
	a "github.com/lindsaygelle/animalcrossing/animals/mouse"
	s "github.com/lindsaygelle/animalcrossing/species/mouse"
)

// Mouse is an Animal Crossing villager type.
type Mouse struct{}

func (v Mouse) Animal() string {
	return (a.Mouse{}).Id()
}

func (v Mouse) Species() string {
	return (s.Mouse{}).Id()
}
