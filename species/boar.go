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

type boar struct{}

func (b boar) Animal() string       { return animals.Boar.Name() }
func (b boar) Class() string        { return classes.Mammalia.Name() }
func (b boar) Conservation() string { return conservations.LeastConcern.Name() }
func (b boar) Domain() string       { return domains.Eukarya.Name() }
func (b boar) Family() string       { return families.Suidae.Name() }
func (b boar) Fictional() bool      { return false }
func (b boar) Genus() string        { return genuses.Sus.Name() }
func (b boar) Kingdom() string      { return kingdoms.Animalia.Name() }
func (b boar) Order() string        { return orders.Artiodactyla.Name() }
func (b boar) Phylum() string       { return phylums.Chordata.Name() }
func (b boar) Species() string      { return species.SScrofa.Name() }

var (
	Boar Species = boar{}
)
