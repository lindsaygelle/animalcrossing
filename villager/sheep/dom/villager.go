package dom

import (
	"github.com/lindsaygelle/animalcrossing/animal/sheep"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	id string = "dom"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Dom.
	Villager = villager.Villager{
		Animal:  sheep.Animal,
		Id:      id,
		Name:    name,
		Phrase:  phrase,
		Special: special}
)
