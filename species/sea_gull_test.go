package species

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestSeaGullClass(t *testing.T) {
	if ok := SeaGull.Class() == aves; !ok {
		t.Fatal("SeaGull.Class() != aves")
	}
}

func TestSeaGullConservation(t *testing.T) {
	if ok := SeaGull.Conservation() == leastConcern; !ok {
		t.Fatal("SeaGull.Conservation() != leastConcern")
	}
}

func TestSeaGullDomain(t *testing.T) {
	if ok := SeaGull.Domain() == eukarya; !ok {
		t.Fatal("SeaGull.Domain() != eukarya")
	}
}

func TestSeaGullFamily(t *testing.T) {
	if ok := SeaGull.Family() == laridae; !ok {
		t.Fatal("SeaGull.Family() != laridae")
	}
}

func TestSeaGullFictional(t *testing.T) {
	if ok := SeaGull.Fictional() == (!fictional); !ok {
		t.Fatal("SeaGull.Fictional() != false")
	}
}

func TestSeaGullGenus(t *testing.T) {
	if ok := SeaGull.Genus() == na; !ok {
		t.Fatal("SeaGull.Genus() != na")
	}
}

func TestSeaGullKingdom(t *testing.T) {
	if ok := SeaGull.Kingdom() == animalia; !ok {
		t.Fatal("SeaGull.Kingdom() != animalia")
	}
}

func TestSeaGullName(t *testing.T) {
	if ok := SeaGull.Name() == animals.SeaGull.Name(); !ok {
		t.Fatalf("SeaGull.Name != %s", animals.SeaGull.Name())
	}
}

func TestSeaGullOrder(t *testing.T) {
	if ok := SeaGull.Order() == charadriiformes; !ok {
		t.Fatal("SeaGull.Order() != charadriiformes")
	}
}

func TestSeaGullPhylum(t *testing.T) {
	if ok := SeaGull.Phylum() == chordata; !ok {
		t.Fatal("SeaGull.Phylum() != chordata")
	}
}

func TestSeaGullSpecies(t *testing.T) {
	if ok := SeaGull.Species() == na; !ok {
		t.Fatal("SeaGull.Species() != na")
	}
}
