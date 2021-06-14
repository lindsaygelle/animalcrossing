package species

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestOwlClass(t *testing.T) {
	if ok := Owl.Class() == aves; !ok {
		t.Fatal("Owl.Class() != aves")
	}
}

func TestOwlConservation(t *testing.T) {
	if ok := Owl.Conservation() == unknown; !ok {
		t.Fatal("Owl.Conservation() != unknown")
	}
}

func TestOwlDomain(t *testing.T) {
	if ok := Owl.Domain() == eukarya; !ok {
		t.Fatal("Owl.Domain() != eukarya")
	}
}

func TestOwlFamily(t *testing.T) {
	if ok := Owl.Family() == na; !ok {
		t.Fatal("Owl.Family() != na")
	}
}

func TestOwlFictional(t *testing.T) {
	if ok := Owl.Fictional() == (!fictional); !ok {
		t.Fatal("Owl.Fictional() != false")
	}
}

func TestOwlGenus(t *testing.T) {
	if ok := Owl.Genus() == na; !ok {
		t.Fatal("Owl.Genus() != na")
	}
}

func TestOwlKingdom(t *testing.T) {
	if ok := Owl.Kingdom() == animalia; !ok {
		t.Fatal("Owl.Kingdom() != animalia")
	}
}

func TestOwlName(t *testing.T) {
	if ok := Owl.Name() != na; !ok {
		t.Fatalf("Owl.Name != %s", animals.Owl.Name())
	}
}

func TestOwlOrder(t *testing.T) {
	if ok := Owl.Order() == strigiformes; !ok {
		t.Fatal("Owl.Order() != strigiformes")
	}
}

func TestOwlPhylum(t *testing.T) {
	if ok := Owl.Phylum() == chordata; !ok {
		t.Fatal("Owl.Phylum() != chordata")
	}
}

func TestOwlSpecies(t *testing.T) {
	if ok := Owl.Species() == na; !ok {
		t.Fatal("Owl.Species() != na")
	}
}
