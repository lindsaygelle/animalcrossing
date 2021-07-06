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

type cat struct{}

func (c cat) Animal() string       { return animals.Cat.Name() }
func (c cat) Class() string        { return classes.Mammalia.Name() }
func (c cat) Conservation() string { return conservations.Domesticated.Name() }
func (c cat) Domain() string       { return domains.Eukarya.Name() }
func (c cat) Family() string       { return families.Felidae.Name() }
func (c cat) Fictional() bool      { return false }
func (c cat) Genus() string        { return genuses.Felis.Name() }
func (c cat) Kingdom() string      { return kingdoms.Animalia.Name() }
func (c cat) Order() string        { return orders.Carnivora.Name() }
func (c cat) Phylum() string       { return phylums.Chordata.Name() }
func (c cat) Species() string      { return species.FCatus.Name() }

var (
	Cat Species = cat{}
)
