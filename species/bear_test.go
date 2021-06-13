package species

import (
	"github.com/lindsaygelle/animalcrossing/animals"
	"testing"
)

func TestBearClass(t *testing.T) {
	if ok := Bear.Class() == mammalia; !ok {
		t.Fatal("Bear.Class() != mammalia")
	}
}

func TestBearConservation(t *testing.T) {
	if ok := Bear.Conservation() == leastConcern; !ok {
		t.Fatal("Bear.Conservation() != leastConcern")
	}
}

func TestBearDomain(t *testing.T) {
	if ok := Bear.Domain() == eukarya; !ok {
		t.Fatal("Bear.Domain() != eukarya")
	}
}

func TestBearFamily(t *testing.T) {
	if ok := Bear.Family() == ursidae; !ok {
		t.Fatal("Bear.Family() != ursidae")
	}
}

func TestBearGenus(t *testing.T) {
	if ok := Bear.Genus() == na; !ok {
		t.Fatal("Bear.Genus() != na")
	}
}

func TestBearKingdom(t *testing.T) {
	if ok := Bear.Kingdom() == animalia; !ok {
		t.Fatal("Bear.Kingdom() != animalia")
	}
}

func TestBearName(t *testing.T) {
	if ok := Bear.Name() != na; !ok {
		t.Fatalf("Bear.Name != %s", animals.Bear.Name())
	}
}

func TestBearOrder(t *testing.T) {
	if ok := Bear.Order() == carnivora; !ok {
		t.Fatal("Bear.Order() != carnivora")
	}
}

func TestBearPhylum(t *testing.T) {
	if ok := Bear.Phylum() == chordata; !ok {
		t.Fatal("Bear.Phylum() != chordata")
	}
}

func TestBearSpecies(t *testing.T) {
	if ok := Bear.Species() == na; !ok {
		t.Fatal("Bear.Species() != na")
	}
}
