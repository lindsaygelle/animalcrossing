package tammi

import (
	"github.com/lindsaygelle/animalcrossing/animal/monkey"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	id string = "tammi"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Tammi.
	Villager = villager.Villager{
		Animal:  monkey.Animal,
		Id:      id,
		Name:    name,
		Phrase:  phrase,
		Special: special}
)
