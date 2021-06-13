package species

import (
	"github.com/lindsaygelle/animalcrossing/animals"
	"testing"
)

func TestOstrichClass(t *testing.T) {
	if ok := Ostrich.Class() == aves; !ok {
		t.Fatal("Ostrich.Class() != aves")
	}
}

func TestOstrichConservation(t *testing.T) {
	if ok := Ostrich.Conservation() == leastConcern; !ok {
		t.Fatal("Ostrich.Conservation() != leastConcern")
	}
}

func TestOstrichDomain(t *testing.T) {
	if ok := Ostrich.Domain() == eukarya; !ok {
		t.Fatal("Ostrich.Domain() != eukarya")
	}
}

func TestOstrichFamily(t *testing.T) {
	if ok := Ostrich.Family() == struthionidae; !ok {
		t.Fatal("Ostrich.Family() != struthionidae")
	}
}

func TestOstrichGenus(t *testing.T) {
	if ok := Ostrich.Genus() == struthio; !ok {
		t.Fatal("Ostrich.Genus() != struthio")
	}
}

func TestOstrichKingdom(t *testing.T) {
	if ok := Ostrich.Kingdom() == animalia; !ok {
		t.Fatal("Ostrich.Kingdom() != animalia")
	}
}

func TestOstrichName(t *testing.T) {
	if ok := Ostrich.Name() != na; !ok {
		t.Fatalf("Ostrich.Name != %s", animals.Ostrich.Name())
	}
}

func TestOstrichOrder(t *testing.T) {
	if ok := Ostrich.Order() == struthioniformes; !ok {
		t.Fatal("Ostrich.Order() != struthioniformes")
	}
}

func TestOstrichPhylum(t *testing.T) {
	if ok := Ostrich.Phylum() == chordata; !ok {
		t.Fatal("Ostrich.Phylum() != chordata")
	}
}

func TestOstrichSpecies(t *testing.T) {
	if ok := Ostrich.Species() == sCamelus; !ok {
		t.Fatal("Ostrich.Species() != sCamelus")
	}
}
