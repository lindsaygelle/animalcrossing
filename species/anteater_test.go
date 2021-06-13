package species

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestAnteaterClass(t *testing.T) {
	if ok := Anteater.Class() == mammalia; !ok {
		t.Fatal("Anteater.Class() != mammalia")
	}
}

func TestAnteaterConservation(t *testing.T) {
	if ok := Anteater.Conservation() == vulnerable; !ok {
		t.Fatal("Anteater.Conservation() != vulnerable")
	}
}

func TestAnteaterDomain(t *testing.T) {
	if ok := Anteater.Domain() == eukarya; !ok {
		t.Fatal("Anteater.Domain() != eukarya")
	}
}

func TestAnteaterFamily(t *testing.T) {
	if ok := Anteater.Family() == na; !ok {
		t.Fatal("Anteater.Family() != na")
	}
}

func TestAnteaterGenus(t *testing.T) {
	if ok := Anteater.Genus() == na; !ok {
		t.Fatal("Anteater.Genus() != na")
	}
}

func TestAnteaterKingdom(t *testing.T) {
	if ok := Anteater.Kingdom() == animalia; !ok {
		t.Fatal("Anteater.Kingdom() != animalia")
	}
}

func TestAnteaterName(t *testing.T) {
	if ok := Anteater.Name() != na; !ok {
		t.Fatalf("Anteater.Name != %s", animals.Anteater.Name())
	}
}

func TestAnteaterOrder(t *testing.T) {
	if ok := Anteater.Order() == pilosa; !ok {
		t.Fatal("Anteater.Order != pilosa")
	}
}

func TestAnteaterPhylum(t *testing.T) {
	if ok := Anteater.Phylum() == chordata; !ok {
		t.Fatal("Anteater.Phylum() != chordata")
	}
}

func TestAnteaterSpecies(t *testing.T) {
	if ok := Anteater.Species() == na; !ok {
		t.Fatal("Anteater.Species() != na")
	}
}
