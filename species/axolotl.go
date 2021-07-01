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

type axolotl struct{}

func (a axolotl) Animal() string       { return animals.Axolotl.Name() }
func (a axolotl) Class() string        { return classes.Amphibia.Name() }
func (a axolotl) Conservation() string { return conservations.CriticallyEndangered.Name() }
func (a axolotl) Domain() string       { return domains.Eukarya.Name() }
func (a axolotl) Family() string       { return families.Ambystomatidae.Name() }
func (a axolotl) Fictional() bool      { return false }
func (a axolotl) Genus() string        { return genuses.Ambystoma.Name() }
func (a axolotl) Kingdom() string      { return kingdoms.Animalia.Name() }
func (a axolotl) Order() string        { return orders.Caudata.Name() }
func (a axolotl) Phylum() string       { return phylums.Chordata.Name() }
func (a axolotl) Species() string      { return species.AMexicanum.Name() }

var (
	Axolotl Species = axolotl{}
)
