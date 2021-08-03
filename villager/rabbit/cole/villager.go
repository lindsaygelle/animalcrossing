package cole

import (
	"github.com/lindsaygelle/animalcrossing/animal/rabbit"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	id string = "cole"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Cole.
	Villager = villager.Villager{
		Animal:  rabbit.Animal,
		Id:      id,
		Name:    name,
		Phrase:  phrase,
		Special: special}
)
