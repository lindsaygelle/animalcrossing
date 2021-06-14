package species

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestPantherClass(t *testing.T) {
	if ok := Panther.Class() == mammalia; !ok {
		t.Fatal("Panther.Class() != mammalia")
	}
}

func TestPantherConservation(t *testing.T) {
	if ok := Panther.Conservation() == unknown; !ok {
		t.Fatal("Panther.Conservation() != unknown")
	}
}

func TestPantherDomain(t *testing.T) {
	if ok := Panther.Domain() == eukarya; !ok {
		t.Fatal("Panther.Domain() != eukarya")
	}
}

func TestPantherFamily(t *testing.T) {
	if ok := Panther.Family() == felidae; !ok {
		t.Fatal("Panther.Family() != felidae")
	}
}

func TestPantherFictional(t *testing.T) {
	if ok := Panther.Fictional() == (!fictional); !ok {
		t.Fatal("Panther.Fictional() != false")
	}
}

func TestPantherGenus(t *testing.T) {
	if ok := Panther.Genus() == panthera; !ok {
		t.Fatal("Panther.Genus() != panthera")
	}
}

func TestPantherKingdom(t *testing.T) {
	if ok := Panther.Kingdom() == animalia; !ok {
		t.Fatal("Panther.Kingdom() != animalia")
	}
}

func TestPantherName(t *testing.T) {
	if ok := Panther.Name() != na; !ok {
		t.Fatalf("Panther.Name != %s", animals.Panther.Name())
	}
}

func TestPantherOrder(t *testing.T) {
	if ok := Panther.Order() == carnivora; !ok {
		t.Fatal("Panther.Order() != carnivora")
	}
}

func TestPantherPhylum(t *testing.T) {
	if ok := Panther.Phylum() == chordata; !ok {
		t.Fatal("Panther.Phylum() != chordata")
	}
}

func TestPantherSpecies(t *testing.T) {
	if ok := Panther.Species() == na; !ok {
		t.Fatal("Panther.Species() != na")
	}
}
