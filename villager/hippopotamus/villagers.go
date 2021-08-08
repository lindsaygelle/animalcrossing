package hippopotamus

import (
	"github.com/lindsaygelle/animalcrossing/villager"
	"github.com/lindsaygelle/animalcrossing/villager/hippopotamus/bertha"
	"github.com/lindsaygelle/animalcrossing/villager/hippopotamus/biff"
	"github.com/lindsaygelle/animalcrossing/villager/hippopotamus/bitty"
	"github.com/lindsaygelle/animalcrossing/villager/hippopotamus/bubbles"
	"github.com/lindsaygelle/animalcrossing/villager/hippopotamus/clara"
	"github.com/lindsaygelle/animalcrossing/villager/hippopotamus/harry"
	"github.com/lindsaygelle/animalcrossing/villager/hippopotamus/hippeux"
	"github.com/lindsaygelle/animalcrossing/villager/hippopotamus/lulu"
	"github.com/lindsaygelle/animalcrossing/villager/hippopotamus/rocco"
	"github.com/lindsaygelle/animalcrossing/villager/hippopotamus/rollo"
)

var (
	Villagers = villager.NewVillagers(
		bertha.Villager,
		biff.Villager,
		bitty.Villager,
		bubbles.Villager,
		clara.Villager,
		harry.Villager,
		hippeux.Villager,
		lulu.Villager,
		rocco.Villager,
		rollo.Villager)
)
