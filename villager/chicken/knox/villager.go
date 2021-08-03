package knox

import (
	"github.com/lindsaygelle/animalcrossing/animal/chicken"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	id string = "knox"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Knox.
	Villager = villager.Villager{
		Animal:  chicken.Animal,
		Id:      id,
		Name:    name,
		Phrase:  phrase,
		Special: special}
)
