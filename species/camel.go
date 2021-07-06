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
)

type camel struct{}

func (c camel) Animal() string       { return animals.Camel.Name() }
func (c camel) Class() string        { return classes.Mammalia.Name() }
func (c camel) Conservation() string { return conservations.Domesticated.Name() }
func (c camel) Domain() string       { return domains.Eukarya.Name() }
func (c camel) Family() string       { return families.Camelidae.Name() }
func (c camel) Fictional() bool      { return false }
func (c camel) Genus() string        { return genuses.Bos.Name() }
func (c camel) Kingdom() string      { return kingdoms.Animalia.Name() }
func (c camel) Order() string        { return orders.Artiodactyla.Name() }
func (c camel) Phylum() string       { return phylums.Chordata.Name() }
func (c camel) Species() string      { return "" }

var (
	Camel Species = camel{}
)
