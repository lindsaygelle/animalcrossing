package species

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestGhostClass(t *testing.T) {
	if ok := Ghost.Class() == na; !ok {
		t.Fatal("Ghost.Class() != na")
	}
}

func TestGhostConservation(t *testing.T) {
	if ok := Ghost.Conservation() == na; !ok {
		t.Fatal("Ghost.Conservation() != na")
	}
}

func TestGhostDomain(t *testing.T) {
	if ok := Ghost.Domain() == na; !ok {
		t.Fatal("Ghost.Domain() != na")
	}
}

func TestGhostFamily(t *testing.T) {
	if ok := Ghost.Family() == na; !ok {
		t.Fatal("Ghost.Family() != na")
	}
}

func TestGhostFictional(t *testing.T) {
	if ok := Ghost.Fictional() == (fictional); !ok {
		t.Fatal("Ghost.Fictional() != true")
	}
}

func TestGhostGenus(t *testing.T) {
	if ok := Ghost.Genus() == na; !ok {
		t.Fatal("Ghost.Genus() != na")
	}
}

func TestGhostKingdom(t *testing.T) {
	if ok := Ghost.Kingdom() == na; !ok {
		t.Fatal("Ghost.Kingdom() != na")
	}
}

func TestGhostName(t *testing.T) {
	if ok := Ghost.Name() != na; !ok {
		t.Fatalf("Ghost.Name != %s", animals.Ghost.Name())
	}
}

func TestGhostOrder(t *testing.T) {
	if ok := Ghost.Order() == na; !ok {
		t.Fatal("Ghost.Order() != na")
	}
}

func TestGhostPhylum(t *testing.T) {
	if ok := Ghost.Phylum() == na; !ok {
		t.Fatal("Ghost.Phylum() != na")
	}
}

func TestGhostSpecies(t *testing.T) {
	if ok := Ghost.Species() == na; !ok {
		t.Fatal("Ghost.Species() != na")
	}
}
