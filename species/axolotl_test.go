package species

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestAxolotlClass(t *testing.T) {
	if ok := Axolotl.Class() == amphibia; !ok {
		t.Fatal("Axolotl.Class() != amphibia")
	}
}

func TestAxolotlConservation(t *testing.T) {
	if ok := Axolotl.Conservation() == criticallyEndangered; !ok {
		t.Fatal("Axolotl.Conservation() != criticallyEndangered")
	}
}

func TestAxolotlDomain(t *testing.T) {
	if ok := Axolotl.Domain() == eukarya; !ok {
		t.Fatal("Axolotl.Domain() != eukarya")
	}
}

func TestAxolotlFamily(t *testing.T) {
	if ok := Axolotl.Family() == ambystomatidae; !ok {
		t.Fatal("Axolotl.Family() != ambystomatidae")
	}
}

func TestAxolotlGenus(t *testing.T) {
	if ok := Axolotl.Genus() == ambystoma; !ok {
		t.Fatal("Axolotl.Genus() != ambystoma")
	}
}

func TestAxolotlKingdom(t *testing.T) {
	if ok := Axolotl.Kingdom() == animalia; !ok {
		t.Fatal("Axolotl.Kingdom() != animalia")
	}
}

func TestAxolotlName(t *testing.T) {
	if ok := Axolotl.Name() != na; !ok {
		t.Fatalf("Axolotl.Name != %s", animals.Axolotl.Name())
	}
}

func TestAxolotlOrder(t *testing.T) {
	if ok := Axolotl.Order() == caudata; !ok {
		t.Fatal("Axolotl.Order() != caudata")
	}
}

func TestAxolotlPhylum(t *testing.T) {
	if ok := Axolotl.Phylum() == chordata; !ok {
		t.Fatal("Axolotl.Phylum() != chordata")
	}
}

func TestAxolotlSpecies(t *testing.T) {
	if ok := Axolotl.Species() == aMexicanum; !ok {
		t.Fatal("Axolotl.Species() != aMexicanum")
	}
}
