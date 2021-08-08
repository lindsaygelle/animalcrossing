package wolf

import (
	"github.com/lindsaygelle/animalcrossing/villager"
	"github.com/lindsaygelle/animalcrossing/villager/wolf/audie"
	"github.com/lindsaygelle/animalcrossing/villager/wolf/chief"
	"github.com/lindsaygelle/animalcrossing/villager/wolf/dobie"
	"github.com/lindsaygelle/animalcrossing/villager/wolf/fang"
	"github.com/lindsaygelle/animalcrossing/villager/wolf/freya"
	"github.com/lindsaygelle/animalcrossing/villager/wolf/kyle"
	"github.com/lindsaygelle/animalcrossing/villager/wolf/lobo"
	"github.com/lindsaygelle/animalcrossing/villager/wolf/skye"
	"github.com/lindsaygelle/animalcrossing/villager/wolf/tarou"
	"github.com/lindsaygelle/animalcrossing/villager/wolf/vivian"
	"github.com/lindsaygelle/animalcrossing/villager/wolf/whitney"
	"github.com/lindsaygelle/animalcrossing/villager/wolf/wlink"
	"github.com/lindsaygelle/animalcrossing/villager/wolf/wolfgang"
)

var (
	Villagers = villager.NewVillagers(
		audie.Villager,
		chief.Villager,
		dobie.Villager,
		fang.Villager,
		freya.Villager,
		kyle.Villager,
		lobo.Villager,
		skye.Villager,
		tarou.Villager,
		vivian.Villager,
		wlink.Villager,
		whitney.Villager,
		wolfgang.Villager)
)
