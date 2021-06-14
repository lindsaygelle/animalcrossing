package species

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestBeaverClass(t *testing.T) {
	if ok := Beaver.Class() == mammalia; !ok {
		t.Fatal("Beaver.Class() != mammalia")
	}
}

func TestBeaverConservation(t *testing.T) {
	if ok := Beaver.Conservation() == leastConcern; !ok {
		t.Fatal("Beaver.Conservation() != leastConcern")
	}
}

func TestBeaverDomain(t *testing.T) {
	if ok := Beaver.Domain() == eukarya; !ok {
		t.Fatal("Beaver.Domain() != eukarya")
	}
}

func TestBeaverFamily(t *testing.T) {
	if ok := Beaver.Family() == castoridae; !ok {
		t.Fatal("Beaver.Family() != castoridae")
	}
}

func TestBeaverFictional(t *testing.T) {
	if ok := Beaver.Fictional() == (!fictional); !ok {
		t.Fatal("Beaver.Fictional() != false")
	}
}

func TestBeaverGenus(t *testing.T) {
	if ok := Beaver.Genus() == castor; !ok {
		t.Fatal("Beaver.Genus() != castor")
	}
}

func TestBeaverKingdom(t *testing.T) {
	if ok := Beaver.Kingdom() == animalia; !ok {
		t.Fatal("Beaver.Kingdom() != animalia")
	}
}

func TestBeaverName(t *testing.T) {
	if ok := Beaver.Name() != na; !ok {
		t.Fatalf("Beaver.Name != %s", animals.Beaver.Name())
	}
}

func TestBeaverOrder(t *testing.T) {
	if ok := Beaver.Order() == rodentia; !ok {
		t.Fatal("Beaver.Order() != rodentia")
	}
}

func TestBeaverPhylum(t *testing.T) {
	if ok := Beaver.Phylum() == chordata; !ok {
		t.Fatal("Beaver.Phylum() != chordata")
	}
}

func TestBeaverSpecies(t *testing.T) {
	if ok := Beaver.Species() == na; !ok {
		t.Fatal("Beaver.Species() != na")
	}
}
