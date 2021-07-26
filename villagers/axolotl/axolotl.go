
package axolotl

import (
	a "github.com/lindsaygelle/animalcrossing/animals/axolotl"
	s "github.com/lindsaygelle/animalcrossing/species/axolotl"
)

// Axolotl is an Animal Crossing villager type.
type Axolotl struct{}

func (v Axolotl) Animal() string {
	return (a.Axolotl{}).Id()
}

func (v Axolotl) Species() string {
	return (s.Axolotl{}).Id()
}
