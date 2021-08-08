package monkey

import (
	"github.com/lindsaygelle/animalcrossing/villager"
	"github.com/lindsaygelle/animalcrossing/villager/monkey/champ"
	"github.com/lindsaygelle/animalcrossing/villager/monkey/deli"
	"github.com/lindsaygelle/animalcrossing/villager/monkey/elise"
	"github.com/lindsaygelle/animalcrossing/villager/monkey/flip"
	"github.com/lindsaygelle/animalcrossing/villager/monkey/monty"
	"github.com/lindsaygelle/animalcrossing/villager/monkey/nana"
	"github.com/lindsaygelle/animalcrossing/villager/monkey/shari"
	"github.com/lindsaygelle/animalcrossing/villager/monkey/simon"
	"github.com/lindsaygelle/animalcrossing/villager/monkey/tammi"
)

var (
	Villagers = villager.NewVillagers(
		champ.Villager,
		deli.Villager,
		elise.Villager,
		flip.Villager,
		monty.Villager,
		nana.Villager,
		shari.Villager,
		simon.Villager,
		tammi.Villager)
)
