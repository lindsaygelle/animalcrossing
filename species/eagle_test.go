package species

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestEagleClass(t *testing.T) {
	if ok := Eagle.Class() == aves; !ok {
		t.Fatal("Eagle.Class() != aves")
	}
}

func TestEagleConservation(t *testing.T) {
	if ok := Eagle.Conservation() == leastConcern; !ok {
		t.Fatal("Eagle.Conservation() != leastConcern")
	}
}

func TestEagleDomain(t *testing.T) {
	if ok := Eagle.Domain() == eukarya; !ok {
		t.Fatal("Eagle.Domain() != eukarya")
	}
}

func TestEagleFamily(t *testing.T) {
	if ok := Eagle.Family() == accipitridae; !ok {
		t.Fatal("Eagle.Family() != accipitridae")
	}
}

func TestEagleFictional(t *testing.T) {
	if ok := Eagle.Fictional() == (!fictional); !ok {
		t.Fatal("Eagle.Fictional() != false")
	}
}

func TestEagleGenus(t *testing.T) {
	if ok := Eagle.Genus() == na; !ok {
		t.Fatal("Eagle.Genus() != na")
	}
}

func TestEagleKingdom(t *testing.T) {
	if ok := Eagle.Kingdom() == animalia; !ok {
		t.Fatal("Eagle.Kingdom() != animalia")
	}
}

func TestEagleName(t *testing.T) {
	if ok := Eagle.Name() != na; !ok {
		t.Fatalf("Eagle.Name != %s", animals.Eagle.Name())
	}
}

func TestEagleOrder(t *testing.T) {
	if ok := Eagle.Order() == accipitriformes; !ok {
		t.Fatal("Eagle.Order() != accipitriformes")
	}
}

func TestEaglePhylum(t *testing.T) {
	if ok := Eagle.Phylum() == chordata; !ok {
		t.Fatal("Eagle.Phylum() != chordata")
	}
}

func TestEagleSpecies(t *testing.T) {
	if ok := Eagle.Species() == na; !ok {
		t.Fatal("Eagle.Species() != na")
	}
}
