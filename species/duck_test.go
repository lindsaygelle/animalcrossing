package species

import (
	"github.com/lindsaygelle/animalcrossing/animals"
	"testing"
)

func TestDuckClass(t *testing.T) {
	if ok := Duck.Class() == aves; !ok {
		t.Fatal("Duck.Class() != aves")
	}
}

func TestDuckConservation(t *testing.T) {
	if ok := Duck.Conservation() == leastConcern; !ok {
		t.Fatal("Duck.Conservation() != leastConcern")
	}
}

func TestDuckDomain(t *testing.T) {
	if ok := Duck.Domain() == eukarya; !ok {
		t.Fatal("Duck.Domain() != eukarya")
	}
}

func TestDuckFamily(t *testing.T) {
	if ok := Duck.Family() == anatidae; !ok {
		t.Fatal("Duck.Family() != anatidae")
	}
}

func TestDuckGenus(t *testing.T) {
	if ok := Duck.Genus() == na; !ok {
		t.Fatal("Duck.Genus() != na")
	}
}

func TestDuckKingdom(t *testing.T) {
	if ok := Duck.Kingdom() == enimalia; !ok {
		t.Fatal("Duck.Kingdom() != enimalia")
	}
}

func TestDuckName(t *testing.T) {
	if ok := Duck.Name() != na; !ok {
		t.Fatalf("Duck.Name != %s", animals.Duck.Name())
	}
}

func TestDuckOrder(t *testing.T) {
	if ok := Duck.Order() == anseriformes; !ok {
		t.Fatal("Duck.Order() != anseriformes")
	}
}

func TestDuckPhylum(t *testing.T) {
	if ok := Duck.Phylum() == chordata; !ok {
		t.Fatal("Duck.Phylum() != chordata")
	}
}

func TestDuckSpecies(t *testing.T) {
	if ok := Duck.Species() == na; !ok {
		t.Fatal("Duck.Species() != na")
	}
}
