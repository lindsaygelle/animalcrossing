package gladys

import (
	"github.com/lindsaygelle/animalcrossing/animal/ostrich"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	id string = "gladys"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Gladys.
	Villager = villager.Villager{
		Animal:  ostrich.Animal,
		Id:      id,
		Name:    name,
		Phrase:  phrase,
		Special: special}
)
