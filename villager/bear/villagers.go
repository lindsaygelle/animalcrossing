package bear

import (
	"github.com/lindsaygelle/animalcrossing/villager"
	"github.com/lindsaygelle/animalcrossing/villager/bear/beardo"
	"github.com/lindsaygelle/animalcrossing/villager/bear/charlise"
	"github.com/lindsaygelle/animalcrossing/villager/bear/chow"
	"github.com/lindsaygelle/animalcrossing/villager/bear/curt"
	"github.com/lindsaygelle/animalcrossing/villager/bear/dozer"
	"github.com/lindsaygelle/animalcrossing/villager/bear/grizzly"
	"github.com/lindsaygelle/animalcrossing/villager/bear/groucho"
	"github.com/lindsaygelle/animalcrossing/villager/bear/ike"
	"github.com/lindsaygelle/animalcrossing/villager/bear/klaus"
	"github.com/lindsaygelle/animalcrossing/villager/bear/megan"
	"github.com/lindsaygelle/animalcrossing/villager/bear/nate"
	"github.com/lindsaygelle/animalcrossing/villager/bear/paula"
	"github.com/lindsaygelle/animalcrossing/villager/bear/pinky"
	"github.com/lindsaygelle/animalcrossing/villager/bear/teddy"
	"github.com/lindsaygelle/animalcrossing/villager/bear/tutu"
	"github.com/lindsaygelle/animalcrossing/villager/bear/ursala"
)

var (
	Villagers = villager.NewVillagers(
		beardo.Villager,
		charlise.Villager,
		chow.Villager,
		curt.Villager,
		dozer.Villager,
		grizzly.Villager,
		groucho.Villager,
		ike.Villager,
		klaus.Villager,
		megan.Villager,
		nate.Villager,
		paula.Villager,
		pinky.Villager,
		teddy.Villager,
		tutu.Villager,
		ursala.Villager)
)
