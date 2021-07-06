package species

import (
	"github.com/lindsaygelle/animalcrossing/animals"
	"github.com/lindsaygelle/animalcrossing/species/classes"
	"github.com/lindsaygelle/animalcrossing/species/conservations"
	"github.com/lindsaygelle/animalcrossing/species/domains"
	"github.com/lindsaygelle/animalcrossing/species/families"
	"github.com/lindsaygelle/animalcrossing/species/genuses"
	"github.com/lindsaygelle/animalcrossing/species/kingdoms"
	"github.com/lindsaygelle/animalcrossing/species/orders"
	"github.com/lindsaygelle/animalcrossing/species/phylums"
	"github.com/lindsaygelle/animalcrossing/species/species"
)

type bovine struct{}

func (b bovine) Animal() string       { return animals.Bovine.Name() }
func (b bovine) Class() string        { return classes.Mammalia.Name() }
func (b bovine) Conservation() string { return conservations.Domesticated.Name() }
func (b bovine) Domain() string       { return domains.Eukarya.Name() }
func (b bovine) Family() string       { return families.Bovidae.Name() }
func (b bovine) Fictional() bool      { return false }
func (b bovine) Genus() string        { return genuses.Bos.Name() }
func (b bovine) Kingdom() string      { return kingdoms.Animalia.Name() }
func (b bovine) Order() string        { return orders.Artiodactyla.Name() }
func (b bovine) Phylum() string       { return phylums.Chordata.Name() }
func (b bovine) Species() string      { return species.BTaurus.Name() }

var (
	Bovine Species = bovine{}
)
