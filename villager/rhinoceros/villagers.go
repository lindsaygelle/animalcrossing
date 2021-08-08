package rhinoceros

import (
	"github.com/lindsaygelle/animalcrossing/villager"
	"github.com/lindsaygelle/animalcrossing/villager/rhinoceros/hornsby"
	"github.com/lindsaygelle/animalcrossing/villager/rhinoceros/merengue"
	"github.com/lindsaygelle/animalcrossing/villager/rhinoceros/patricia"
	"github.com/lindsaygelle/animalcrossing/villager/rhinoceros/petunia"
	"github.com/lindsaygelle/animalcrossing/villager/rhinoceros/renee"
	"github.com/lindsaygelle/animalcrossing/villager/rhinoceros/rhonda"
	"github.com/lindsaygelle/animalcrossing/villager/rhinoceros/spike"
	"github.com/lindsaygelle/animalcrossing/villager/rhinoceros/tank"
	"github.com/lindsaygelle/animalcrossing/villager/rhinoceros/tiara"
)

var (
	Villagers = villager.NewVillagers(
		hornsby.Villager,
		merengue.Villager,
		patricia.Villager,
		petunia.Villager,
		renee.Villager,
		rhonda.Villager,
		spike.Villager,
		tank.Villager,
		tiara.Villager)
)
