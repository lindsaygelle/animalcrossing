package lopez

import (
	"github.com/lindsaygelle/animalcrossing/animal/deer"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	id string = "lopez"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Lopez.
	Villager = villager.Villager{
		Animal:  deer.Animal,
		Id:      id,
		Name:    name,
		Phrase:  phrase,
		Special: special}
)
