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

type chicken struct{}

func (c chicken) Animal() string       { return animals.Chicken.Name() }
func (c chicken) Class() string        { return classes.Aves.Name() }
func (c chicken) Conservation() string { return conservations.Domesticated.Name() }
func (c chicken) Domain() string       { return domains.Eukarya.Name() }
func (c chicken) Family() string       { return families.Phasianidae.Name() }
func (c chicken) Fictional() bool      { return false }
func (c chicken) Genus() string        { return genuses.Gallus.Name() }
func (c chicken) Kingdom() string      { return kingdoms.Animalia.Name() }
func (c chicken) Order() string        { return orders.Galliformes.Name() }
func (c chicken) Phylum() string       { return phylums.Chordata.Name() }
func (c chicken) Species() string      { return species.GGallus.Name() }

var (
	Chicken Species = chicken{}
)
