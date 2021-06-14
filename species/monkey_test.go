package species

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestMonkeyClass(t *testing.T) {
	if ok := Monkey.Class() == mammalia; !ok {
		t.Fatal("Monkey.Class() != mammalia")
	}
}

func TestMonkeyConservation(t *testing.T) {
	if ok := Monkey.Conservation() == endangered; !ok {
		t.Fatal("Monkey.Conservation() != endangered")
	}
}

func TestMonkeyDomain(t *testing.T) {
	if ok := Monkey.Domain() == eukarya; !ok {
		t.Fatal("Monkey.Domain() != eukarya")
	}
}

func TestMonkeyFamily(t *testing.T) {
	if ok := Monkey.Family() == na; !ok {
		t.Fatal("Monkey.Family() != na")
	}
}

func TestMonkeyFictional(t *testing.T) {
	if ok := Monkey.Fictional() == (!fictional); !ok {
		t.Fatal("Monkey.Fictional() != false")
	}
}

func TestMonkeyGenus(t *testing.T) {
	if ok := Monkey.Genus() == na; !ok {
		t.Fatal("Monkey.Genus() != na")
	}
}

func TestMonkeyKingdom(t *testing.T) {
	if ok := Monkey.Kingdom() == animalia; !ok {
		t.Fatal("Monkey.Kingdom() != animalia")
	}
}

func TestMonkeyName(t *testing.T) {
	if ok := Monkey.Name() != na; !ok {
		t.Fatalf("Monkey.Name != %s", animals.Monkey.Name())
	}
}

func TestMonkeyOrder(t *testing.T) {
	if ok := Monkey.Order() == primates; !ok {
		t.Fatal("Monkey.Order() != primates")
	}
}

func TestMonkeyPhylum(t *testing.T) {
	if ok := Monkey.Phylum() == chordata; !ok {
		t.Fatal("Monkey.Phylum() != chordata")
	}
}

func TestMonkeySpecies(t *testing.T) {
	if ok := Monkey.Species() == na; !ok {
		t.Fatal("Monkey.Species() != na")
	}
}
