package antonio

import (
	"github.com/lindsaygelle/animalcrossing/animal/anteater"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	id string = "antonio"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Antonio.
	Villager = villager.Villager{
		Animal:  anteater.Animal,
		Id:      id,
		Name:    name,
		Phrase:  phrase,
		Special: special}
)
