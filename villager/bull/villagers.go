package bull

import (
	"github.com/lindsaygelle/animalcrossing/villager"
	"github.com/lindsaygelle/animalcrossing/villager/bull/angus"
	"github.com/lindsaygelle/animalcrossing/villager/bull/chuck"
	"github.com/lindsaygelle/animalcrossing/villager/bull/coach"
	"github.com/lindsaygelle/animalcrossing/villager/bull/oxford"
	"github.com/lindsaygelle/animalcrossing/villager/bull/rodeo"
	"github.com/lindsaygelle/animalcrossing/villager/bull/stu"
	"github.com/lindsaygelle/animalcrossing/villager/bull/tbone"
	"github.com/lindsaygelle/animalcrossing/villager/bull/vic"
	"github.com/lindsaygelle/animalcrossing/villager/bull/weldon"
)

var (
	Villagers = villager.NewVillagers(
		angus.Villager,
		chuck.Villager,
		coach.Villager,
		oxford.Villager,
		rodeo.Villager,
		stu.Villager,
		tbone.Villager,
		vic.Villager,
		weldon.Villager)
)
