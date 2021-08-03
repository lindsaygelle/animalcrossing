package graham

import (
	"github.com/lindsaygelle/animalcrossing/animal/hamster"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	id string = "graham"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Graham.
	Villager = villager.Villager{
		Animal:  hamster.Animal,
		Id:      id,
		Name:    name,
		Phrase:  phrase,
		Special: special}
)
