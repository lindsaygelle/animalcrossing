package klaus

import (
	"github.com/lindsaygelle/animalcrossing/animal/bear"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	id string = "klaus"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Klaus.
	Villager = villager.Villager{
		Animal:  bear.Animal,
		Id:      id,
		Name:    name,
		Phrase:  phrase,
		Special: special}
)
