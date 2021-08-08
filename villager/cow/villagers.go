package cow

import (
	"github.com/lindsaygelle/animalcrossing/villager"
	"github.com/lindsaygelle/animalcrossing/villager/cow/belle"
	"github.com/lindsaygelle/animalcrossing/villager/cow/bessie"
	"github.com/lindsaygelle/animalcrossing/villager/cow/carrot"
	"github.com/lindsaygelle/animalcrossing/villager/cow/naomi"
	"github.com/lindsaygelle/animalcrossing/villager/cow/norma"
	"github.com/lindsaygelle/animalcrossing/villager/cow/patty"
	"github.com/lindsaygelle/animalcrossing/villager/cow/petunia"
	"github.com/lindsaygelle/animalcrossing/villager/cow/tipper"
)

var (
	Villagers = villager.NewVillagers(
		belle.Villager,
		bessie.Villager,
		carrot.Villager,
		naomi.Villager,
		norma.Villager,
		patty.Villager,
		petunia.Villager,
		tipper.Villager)
)
