package vic

import (
	"github.com/lindsaygelle/animalcrossing/animal/bull"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	id string = "vic"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Vic.
	Villager = villager.Villager{
		Animal:  bull.Animal,
		Id:      id,
		Name:    name,
		Phrase:  phrase,
		Special: special}
)
