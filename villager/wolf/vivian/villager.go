package vivian

import (
	"github.com/lindsaygelle/animalcrossing/animal/wolf"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	id string = "vivian"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Vivian.
	Villager = villager.Villager{
		Animal:  wolf.Animal,
		Id:      id,
		Name:    name,
		Phrase:  phrase,
		Special: special}
)
