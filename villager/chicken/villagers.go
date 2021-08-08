package chicken

import (
	"github.com/lindsaygelle/animalcrossing/villager"
	"github.com/lindsaygelle/animalcrossing/villager/chicken/ava"
	"github.com/lindsaygelle/animalcrossing/villager/chicken/becky"
	"github.com/lindsaygelle/animalcrossing/villager/chicken/benedict"
	"github.com/lindsaygelle/animalcrossing/villager/chicken/betty"
	"github.com/lindsaygelle/animalcrossing/villager/chicken/broffina"
	"github.com/lindsaygelle/animalcrossing/villager/chicken/egbert"
	"github.com/lindsaygelle/animalcrossing/villager/chicken/goose"
	"github.com/lindsaygelle/animalcrossing/villager/chicken/hank"
	"github.com/lindsaygelle/animalcrossing/villager/chicken/hector"
	"github.com/lindsaygelle/animalcrossing/villager/chicken/ken"
	"github.com/lindsaygelle/animalcrossing/villager/chicken/knox"
	"github.com/lindsaygelle/animalcrossing/villager/chicken/leigh"
	"github.com/lindsaygelle/animalcrossing/villager/chicken/plucky"
	"github.com/lindsaygelle/animalcrossing/villager/chicken/rhoda"
)

var (
	Villagers = villager.NewVillagers(
		ava.Villager,
		becky.Villager,
		benedict.Villager,
		betty.Villager,
		broffina.Villager,
		egbert.Villager,
		goose.Villager,
		hank.Villager,
		hector.Villager,
		ken.Villager,
		knox.Villager,
		leigh.Villager,
		plucky.Villager,
		rhoda.Villager)
)
