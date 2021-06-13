package species

import (
	"github.com/lindsaygelle/animalcrossing/animals"
	"testing"
)

func TestGoatClass(t *testing.T) {
	if ok := Goat.Class() == mammalia; !ok {
		t.Fatal("Goat.Class() != mammalia")
	}
}

func TestGoatConservation(t *testing.T) {
	if ok := Goat.Conservation() == domesticated; !ok {
		t.Fatal("Goat.Conservation() != domesticated")
	}
}

func TestGoatDomain(t *testing.T) {
	if ok := Goat.Domain() == eukarya; !ok {
		t.Fatal("Goat.Domain() != eukarya")
	}
}

func TestGoatFamily(t *testing.T) {
	if ok := Goat.Family() == bovidae; !ok {
		t.Fatal("Goat.Family() != bovidae")
	}
}

func TestGoatGenus(t *testing.T) {
	if ok := Goat.Genus() == capra; !ok {
		t.Fatal("Goat.Genus() != capra")
	}
}

func TestGoatKingdom(t *testing.T) {
	if ok := Goat.Kingdom() == animalia; !ok {
		t.Fatal("Goat.Kingdom() != animalia")
	}
}

func TestGoatName(t *testing.T) {
	if ok := Goat.Name() != na; !ok {
		t.Fatalf("Goat.Name != %s", animals.Goat.Name())
	}
}

func TestGoatOrder(t *testing.T) {
	if ok := Goat.Order() == artiodactyla; !ok {
		t.Fatal("Goat.Order() != artiodactyla")
	}
}

func TestGoatPhylum(t *testing.T) {
	if ok := Goat.Phylum() == chordata; !ok {
		t.Fatal("Goat.Phylum() != chordata")
	}
}

func TestGoatSpecies(t *testing.T) {
	if ok := Goat.Species() == cAegagrus; !ok {
		t.Fatal("Goat.Species() != cAegagrus")
	}
}
