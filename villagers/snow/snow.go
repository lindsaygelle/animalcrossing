
package snow

import (
	a "github.com/lindsaygelle/animalcrossing/animals/snow"
	s "github.com/lindsaygelle/animalcrossing/species/snow"
)

// Snow is an Animal Crossing villager type.
type Snow struct{}

func (v Snow) Animal() string {
	return (a.Snow{}).Id()
}

func (v Snow) Species() string {
	return (s.Snow{}).Id()
}
