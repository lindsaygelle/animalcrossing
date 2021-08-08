package hamster

import (
	"github.com/lindsaygelle/animalcrossing/villager"
	"github.com/lindsaygelle/animalcrossing/villager/hamster/apple"
	"github.com/lindsaygelle/animalcrossing/villager/hamster/clay"
	"github.com/lindsaygelle/animalcrossing/villager/hamster/flurry"
	"github.com/lindsaygelle/animalcrossing/villager/hamster/graham"
	"github.com/lindsaygelle/animalcrossing/villager/hamster/hamlet"
	"github.com/lindsaygelle/animalcrossing/villager/hamster/hamphrey"
	"github.com/lindsaygelle/animalcrossing/villager/hamster/holden"
	"github.com/lindsaygelle/animalcrossing/villager/hamster/rodney"
	"github.com/lindsaygelle/animalcrossing/villager/hamster/soleil"
)

var (
	Villagers = villager.NewVillagers(
		apple.Villager,
		clay.Villager,
		flurry.Villager,
		graham.Villager,
		hamlet.Villager,
		hamphrey.Villager,
		holden.Villager,
		rodney.Villager,
		soleil.Villager)
)
