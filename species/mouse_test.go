package species

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestMouseClass(t *testing.T) {
	if ok := Mouse.Class() == mammalia; !ok {
		t.Fatal("Mouse.Class() != mammalia")
	}
}

func TestMouseConservation(t *testing.T) {
	if ok := Mouse.Conservation() == leastConcern; !ok {
		t.Fatal("Mouse.Conservation() != leastConcern")
	}
}

func TestMouseDomain(t *testing.T) {
	if ok := Mouse.Domain() == eukarya; !ok {
		t.Fatal("Mouse.Domain() != eukarya")
	}
}

func TestMouseFamily(t *testing.T) {
	if ok := Mouse.Family() == muridae; !ok {
		t.Fatal("Mouse.Family() != muridae")
	}
}

func TestMouseFictional(t *testing.T) {
	if ok := Mouse.Fictional() == (!fictional); !ok {
		t.Fatal("Mouse.Fictional() != false")
	}
}

func TestMouseGenus(t *testing.T) {
	if ok := Mouse.Genus() == mus; !ok {
		t.Fatal("Mouse.Genus() != mus")
	}
}

func TestMouseKingdom(t *testing.T) {
	if ok := Mouse.Kingdom() == animalia; !ok {
		t.Fatal("Mouse.Kingdom() != animalia")
	}
}

func TestMouseName(t *testing.T) {
	if ok := Mouse.Name() != na; !ok {
		t.Fatalf("Mouse.Name != %s", animals.Mouse.Name())
	}
}

func TestMouseOrder(t *testing.T) {
	if ok := Mouse.Order() == rodentia; !ok {
		t.Fatal("Mouse.Order() != rodentia")
	}
}

func TestMousePhylum(t *testing.T) {
	if ok := Mouse.Phylum() == chordata; !ok {
		t.Fatal("Mouse.Phylum() != chordata")
	}
}

func TestMouseSpecies(t *testing.T) {
	if ok := Mouse.Species() == na; !ok {
		t.Fatal("Mouse.Species() != na")
	}
}
