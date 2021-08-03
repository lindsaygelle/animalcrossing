package leonardo

import (
	"github.com/lindsaygelle/animalcrossing/animal/tiger"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	id string = "leonardo"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Leonardo.
	Villager = villager.Villager{
		Animal:  tiger.Animal,
		Id:      id,
		Name:    name,
		Phrase:  phrase,
		Special: special}
)
