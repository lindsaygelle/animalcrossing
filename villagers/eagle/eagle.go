
package eagle

import (
	a "github.com/lindsaygelle/animalcrossing/animals/eagle"
	s "github.com/lindsaygelle/animalcrossing/species/eagle"
)

// Eagle is an Animal Crossing villager type.
type Eagle struct{}

func (v Eagle) Animal() string {
	return (a.Eagle{}).Id()
}

func (v Eagle) Species() string {
	return (s.Eagle{}).Id()
}
