package species

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestWolfClass(t *testing.T) {
	if ok := Wolf.Class() == mammalia; !ok {
		t.Fatal("Wolf.Class() != mammalia")
	}
}

func TestWolfConservation(t *testing.T) {
	if ok := Wolf.Conservation() == leastConcern; !ok {
		t.Fatal("Wolf.Conservation() != leastConcern")
	}
}

func TestWolfDomain(t *testing.T) {
	if ok := Wolf.Domain() == eukarya; !ok {
		t.Fatal("Wolf.Domain() != eukarya")
	}
}

func TestWolfFamily(t *testing.T) {
	if ok := Wolf.Family() == canidae; !ok {
		t.Fatal("Wolf.Family() != canidae")
	}
}

func TestWolfFictional(t *testing.T) {
	if ok := Wolf.Fictional() == (!fictional); !ok {
		t.Fatal("Wolf.Fictional() != false")
	}
}

func TestWolfGenus(t *testing.T) {
	if ok := Wolf.Genus() == canis; !ok {
		t.Fatal("Wolf.Genus() != canis")
	}
}

func TestWolfKingdom(t *testing.T) {
	if ok := Wolf.Kingdom() == animalia; !ok {
		t.Fatal("Wolf.Kingdom() != animalia")
	}
}

func TestWolfName(t *testing.T) {
	if ok := Wolf.Name() != na; !ok {
		t.Fatalf("Wolf.Name != %s", animals.Wolf.Name())
	}
}

func TestWolfOrder(t *testing.T) {
	if ok := Wolf.Order() == carnivora; !ok {
		t.Fatal("Wolf.Order() != carnivora")
	}
}

func TestWolfPhylum(t *testing.T) {
	if ok := Wolf.Phylum() == chordata; !ok {
		t.Fatal("Wolf.Phylum() != chordata")
	}
}

func TestWolfSpecies(t *testing.T) {
	if ok := Wolf.Species() == na; !ok {
		t.Fatal("Wolf.Species() != na")
	}
}
