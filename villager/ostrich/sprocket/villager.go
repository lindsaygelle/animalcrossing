package sprocket

import (
	"github.com/lindsaygelle/animalcrossing/animal/ostrich"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	id string = "sprocket"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Sprocket.
	Villager = villager.Villager{
		Animal:  ostrich.Animal,
		Id:      id,
		Name:    name,
		Phrase:  phrase,
		Special: special}
)
