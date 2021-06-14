package species

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestOctopusClass(t *testing.T) {
	if ok := Octopus.Class() == cephalopoda; !ok {
		t.Fatal("Octopus.Class() != cephalopoda")
	}
}

func TestOctopusConservation(t *testing.T) {
	if ok := Octopus.Conservation() == leastConcern; !ok {
		t.Fatal("Octopus.Conservation() != leastConcern")
	}
}

func TestOctopusDomain(t *testing.T) {
	if ok := Octopus.Domain() == eukarya; !ok {
		t.Fatal("Octopus.Domain() != eukarya")
	}
}

func TestOctopusFamily(t *testing.T) {
	if ok := Octopus.Family() == na; !ok {
		t.Fatal("Octopus.Family() != na")
	}
}

func TestOctopusFictional(t *testing.T) {
	if ok := Octopus.Fictional() == (!fictional); !ok {
		t.Fatal("Octopus.Fictional() != false")
	}
}

func TestOctopusGenus(t *testing.T) {
	if ok := Octopus.Genus() == na; !ok {
		t.Fatal("Octopus.Genus() != na")
	}
}

func TestOctopusKingdom(t *testing.T) {
	if ok := Octopus.Kingdom() == animalia; !ok {
		t.Fatal("Octopus.Kingdom() != animalia")
	}
}

func TestOctopusName(t *testing.T) {
	if ok := Octopus.Name() != na; !ok {
		t.Fatalf("Octopus.Name != %s", animals.Octopus.Name())
	}
}

func TestOctopusOrder(t *testing.T) {
	if ok := Octopus.Order() == octopoda; !ok {
		t.Fatal("Octopus.Order() != octopoda")
	}
}

func TestOctopusPhylum(t *testing.T) {
	if ok := Octopus.Phylum() == mollusca; !ok {
		t.Fatal("Octopus.Phylum() != mollusca")
	}
}

func TestOctopusSpecies(t *testing.T) {
	if ok := Octopus.Species() == na; !ok {
		t.Fatal("Octopus.Species() != na")
	}
}
