package boris

import (
	"github.com/lindsaygelle/animalcrossing/animal/pig"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	id string = "boris"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Boris.
	Villager = villager.Villager{
		Animal:  pig.Animal,
		Id:      id,
		Name:    name,
		Phrase:  phrase,
		Special: special}
)
