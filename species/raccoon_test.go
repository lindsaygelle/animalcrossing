package species

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestRaccoonClass(t *testing.T) {
	if ok := Raccoon.Class() == mammalia; !ok {
		t.Fatal("Raccoon.Class() != mammalia")
	}
}

func TestRaccoonConservation(t *testing.T) {
	if ok := Raccoon.Conservation() == leastConcern; !ok {
		t.Fatal("Raccoon.Conservation() != leastConcern")
	}
}

func TestRaccoonDomain(t *testing.T) {
	if ok := Raccoon.Domain() == eukarya; !ok {
		t.Fatal("Raccoon.Domain() != eukarya")
	}
}

func TestRaccoonFamily(t *testing.T) {
	if ok := Raccoon.Family() == canidae; !ok {
		t.Fatal("Raccoon.Family() != canidae")
	}
}

func TestRaccoonFictional(t *testing.T) {
	if ok := Raccoon.Fictional() == (!fictional); !ok {
		t.Fatal("Raccoon.Fictional() != false")
	}
}

func TestRaccoonGenus(t *testing.T) {
	if ok := Raccoon.Genus() == nyctereutes; !ok {
		t.Fatal("Raccoon.Genus() != nyctereutes")
	}
}

func TestRaccoonKingdom(t *testing.T) {
	if ok := Raccoon.Kingdom() == animalia; !ok {
		t.Fatal("Raccoon.Kingdom() != animalia")
	}
}

func TestRaccoonName(t *testing.T) {
	if ok := Raccoon.Name() != na; !ok {
		t.Fatalf("Raccoon.Name != %s", animals.Raccoon.Name())
	}
}

func TestRaccoonOrder(t *testing.T) {
	if ok := Raccoon.Order() == carnivora; !ok {
		t.Fatal("Raccoon.Order() != carnivora")
	}
}

func TestRaccoonPhylum(t *testing.T) {
	if ok := Raccoon.Phylum() == chordata; !ok {
		t.Fatal("Raccoon.Phylum() != chordata")
	}
}

func TestRaccoonSpecies(t *testing.T) {
	if ok := Raccoon.Species() == nProcyonoides; !ok {
		t.Fatal("Raccoon.Species() != nProcyonoides")
	}
}
