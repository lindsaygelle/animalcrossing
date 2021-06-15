package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	tBone string = "T-Bone"
)

var (
	// TBone is an Animal Crossing villager.
	//
	// TBone is a Bull.
	TBone Villager = villager{
		animal: animals.Bull.Name(),
		name:   tBone}
)
