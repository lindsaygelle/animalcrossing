package sandy

import (
	"github.com/lindsaygelle/animalcrossing/animal/ostrich"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	id string = "sandy"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Sandy.
	Villager = villager.Villager{
		Animal:  ostrich.Animal,
		Id:      id,
		Name:    name,
		Phrase:  phrase,
		Special: special}
)
