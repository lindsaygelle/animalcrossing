package penguin

import (
	"github.com/lindsaygelle/animalcrossing/villager"
	"github.com/lindsaygelle/animalcrossing/villager/penguin/analog"
	"github.com/lindsaygelle/animalcrossing/villager/penguin/aurora"
	"github.com/lindsaygelle/animalcrossing/villager/penguin/boomer"
	"github.com/lindsaygelle/animalcrossing/villager/penguin/cube"
	"github.com/lindsaygelle/animalcrossing/villager/penguin/flo"
	"github.com/lindsaygelle/animalcrossing/villager/penguin/friga"
	"github.com/lindsaygelle/animalcrossing/villager/penguin/gwen"
	"github.com/lindsaygelle/animalcrossing/villager/penguin/hopper"
	"github.com/lindsaygelle/animalcrossing/villager/penguin/iggly"
	"github.com/lindsaygelle/animalcrossing/villager/penguin/nobuo"
	"github.com/lindsaygelle/animalcrossing/villager/penguin/puck"
	"github.com/lindsaygelle/animalcrossing/villager/penguin/roald"
	"github.com/lindsaygelle/animalcrossing/villager/penguin/sprinkle"
	"github.com/lindsaygelle/animalcrossing/villager/penguin/tex"
	"github.com/lindsaygelle/animalcrossing/villager/penguin/wade"
)

var (
	Villagers = villager.NewVillagers(
		analog.Villager,
		aurora.Villager,
		boomer.Villager,
		cube.Villager,
		flo.Villager,
		friga.Villager,
		gwen.Villager,
		hopper.Villager,
		iggly.Villager,
		nobuo.Villager,
		puck.Villager,
		roald.Villager,
		sprinkle.Villager,
		tex.Villager,
		wade.Villager)
)
