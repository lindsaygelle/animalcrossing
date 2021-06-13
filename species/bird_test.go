package species

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestBirdClass(t *testing.T) {
	if ok := Bird.Class() == aves; !ok {
		t.Fatal("Bird.Class() != aves")
	}
}

func TestBirdConservation(t *testing.T) {
	if ok := Bird.Conservation() == leastConcern; !ok {
		t.Fatal("Bird.Conservation() != leastConcern")
	}
}

func TestBirdDomain(t *testing.T) {
	if ok := Bird.Domain() == eukarya; !ok {
		t.Fatal("Bird.Domain() != eukarya")
	}
}

func TestBirdFamily(t *testing.T) {
	if ok := Bird.Family() == na; !ok {
		t.Fatal("Bird.Family() != na")
	}
}

func TestBirdGenus(t *testing.T) {
	if ok := Bird.Genus() == na; !ok {
		t.Fatal("Bird.Genus() != na")
	}
}

func TestBirdKingdom(t *testing.T) {
	if ok := Bird.Kingdom() == animalia; !ok {
		t.Fatal("Bird.Kingdom() != animalia")
	}
}

func TestBirdName(t *testing.T) {
	if ok := Bird.Name() != na; !ok {
		t.Fatalf("Bird.Name != %s", animals.Bird.Name())
	}
}

func TestBirdOrder(t *testing.T) {
	if ok := Bird.Order() == na; !ok {
		t.Fatal("Bird.Order() != na")
	}
}

func TestBirdPhylum(t *testing.T) {
	if ok := Bird.Phylum() == chordata; !ok {
		t.Fatal("Bird.Phylum() != chordata")
	}
}

func TestBirdSpecies(t *testing.T) {
	if ok := Bird.Species() == na; !ok {
		t.Fatal("Bird.Species() != na")
	}
}
