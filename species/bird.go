package species

import (
	"github.com/lindsaygelle/animalcrossing/animals"
	"github.com/lindsaygelle/animalcrossing/species/classes"
	"github.com/lindsaygelle/animalcrossing/species/conservations"
	"github.com/lindsaygelle/animalcrossing/species/domains"
	"github.com/lindsaygelle/animalcrossing/species/kingdoms"
	"github.com/lindsaygelle/animalcrossing/species/orders"
	"github.com/lindsaygelle/animalcrossing/species/phylums"
)

type bird struct{}

func (b bird) Animal() string       { return animals.Bird.Name() }
func (b bird) Class() string        { return classes.Aves.Name() }
func (b bird) Conservation() string { return conservations.LeastConcern.Name() }
func (b bird) Domain() string       { return domains.Eukarya.Name() }
func (b bird) Family() string       { return "" }
func (b bird) Fictional() bool      { return false }
func (b bird) Genus() string        { return "" }
func (b bird) Kingdom() string      { return kingdoms.Animalia.Name() }
func (b bird) Order() string        { return orders.Rodentia.Name() }
func (b bird) Phylum() string       { return phylums.Chordata.Name() }
func (b bird) Species() string      { return "" }

var (
	Bird Species = bird{}
)
