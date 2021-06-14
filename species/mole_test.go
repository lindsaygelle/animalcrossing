package species

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestMoleClass(t *testing.T) {
	if ok := Mole.Class() == mammalia; !ok {
		t.Fatal("Mole.Class() != mammalia")
	}
}

func TestMoleConservation(t *testing.T) {
	if ok := Mole.Conservation() == leastConcern; !ok {
		t.Fatal("Mole.Conservation() != leastConcern")
	}
}

func TestMoleDomain(t *testing.T) {
	if ok := Mole.Domain() == eukarya; !ok {
		t.Fatal("Mole.Domain() != eukarya")
	}
}

func TestMoleFamily(t *testing.T) {
	if ok := Mole.Family() == talpidae; !ok {
		t.Fatal("Mole.Family() != talpidae")
	}
}

func TestMoleFictional(t *testing.T) {
	if ok := Mole.Fictional() == (!fictional); !ok {
		t.Fatal("Mole.Fictional() != false")
	}
}

func TestMoleGenus(t *testing.T) {
	if ok := Mole.Genus() == na; !ok {
		t.Fatal("Mole.Genus() != na")
	}
}

func TestMoleKingdom(t *testing.T) {
	if ok := Mole.Kingdom() == animalia; !ok {
		t.Fatal("Mole.Kingdom() != animalia")
	}
}

func TestMoleName(t *testing.T) {
	if ok := Mole.Name() != na; !ok {
		t.Fatalf("Mole.Name != %s", animals.Mole.Name())
	}
}

func TestMoleOrder(t *testing.T) {
	if ok := Mole.Order() == eulipotyphla; !ok {
		t.Fatal("Mole.Order() != eulipotyphla")
	}
}

func TestMolePhylum(t *testing.T) {
	if ok := Mole.Phylum() == chordata; !ok {
		t.Fatal("Mole.Phylum() != chordata")
	}
}

func TestMoleSpecies(t *testing.T) {
	if ok := Mole.Species() == na; !ok {
		t.Fatal("Mole.Species() != na")
	}
}
