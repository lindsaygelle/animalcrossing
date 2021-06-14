package species

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestFrogClass(t *testing.T) {
	if ok := Frog.Class() == amphibia; !ok {
		t.Fatal("Frog.Class() != amphibia")
	}
}

func TestFrogConservation(t *testing.T) {
	if ok := Frog.Conservation() == vulnerable; !ok {
		t.Fatal("Frog.Conservation() != vulnerable")
	}
}

func TestFrogDomain(t *testing.T) {
	if ok := Frog.Domain() == eukarya; !ok {
		t.Fatal("Frog.Domain() != eukarya")
	}
}

func TestFrogFamily(t *testing.T) {
	if ok := Frog.Family() == na; !ok {
		t.Fatal("Frog.Family() != na")
	}
}

func TestFrogFictional(t *testing.T) {
	if ok := Frog.Fictional() == (!fictional); !ok {
		t.Fatal("Frog.Fictional() != false")
	}
}

func TestFrogGenus(t *testing.T) {
	if ok := Frog.Genus() == na; !ok {
		t.Fatal("Frog.Genus() != na")
	}
}

func TestFrogKingdom(t *testing.T) {
	if ok := Frog.Kingdom() == animalia; !ok {
		t.Fatal("Frog.Kingdom() != animalia")
	}
}

func TestFrogName(t *testing.T) {
	if ok := Frog.Name() != na; !ok {
		t.Fatalf("Frog.Name != %s", animals.Frog.Name())
	}
}

func TestFrogOrder(t *testing.T) {
	if ok := Frog.Order() == anura; !ok {
		t.Fatal("Frog.Order() != anura")
	}
}

func TestFrogPhylum(t *testing.T) {
	if ok := Frog.Phylum() == chordata; !ok {
		t.Fatal("Frog.Phylum() != chordata")
	}
}

func TestFrogSpecies(t *testing.T) {
	if ok := Frog.Species() == na; !ok {
		t.Fatal("Frog.Species() != na")
	}
}
