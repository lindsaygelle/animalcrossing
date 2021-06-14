package species

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestSquirrelClass(t *testing.T) {
	if ok := Squirrel.Class() == mammalia; !ok {
		t.Fatal("Squirrel.Class() != mammalia")
	}
}

func TestSquirrelConservation(t *testing.T) {
	if ok := Squirrel.Conservation() == leastConcern; !ok {
		t.Fatal("Squirrel.Conservation() != leastConcern")
	}
}

func TestSquirrelDomain(t *testing.T) {
	if ok := Squirrel.Domain() == eukarya; !ok {
		t.Fatal("Squirrel.Domain() != eukarya")
	}
}

func TestSquirrelFamily(t *testing.T) {
	if ok := Squirrel.Family() == sciuridae; !ok {
		t.Fatal("Squirrel.Family() != sciuridae")
	}
}

func TestSquirrelGenus(t *testing.T) {
	if ok := Squirrel.Genus() == na; !ok {
		t.Fatal("Squirrel.Genus() != na")
	}
}

func TestSquirrelKingdom(t *testing.T) {
	if ok := Squirrel.Kingdom() == animalia; !ok {
		t.Fatal("Squirrel.Kingdom() != animalia")
	}
}

func TestSquirrelName(t *testing.T) {
	if ok := Squirrel.Name() != na; !ok {
		t.Fatalf("Squirrel.Name != %s", animals.Squirrel.Name())
	}
}

func TestSquirrelOrder(t *testing.T) {
	if ok := Squirrel.Order() == rodentia; !ok {
		t.Fatal("Squirrel.Order() != rodentia")
	}
}

func TestSquirrelPhylum(t *testing.T) {
	if ok := Squirrel.Phylum() == chordata; !ok {
		t.Fatal("Squirrel.Phylum() != chordata")
	}
}

func TestSquirrelSpecies(t *testing.T) {
	if ok := Squirrel.Species() == na; !ok {
		t.Fatal("Squirrel.Species() != na")
	}
}
