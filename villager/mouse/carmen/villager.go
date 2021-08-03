package carmen

import (
	"github.com/lindsaygelle/animalcrossing/animal/mouse"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	id string = "carmen"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Carmen.
	Villager = villager.Villager{
		Animal:  mouse.Animal,
		Id:      id,
		Name:    name,
		Phrase:  phrase,
		Special: special}
)
