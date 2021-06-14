package species

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestPenguinClass(t *testing.T) {
	if ok := Penguin.Class() == aves; !ok {
		t.Fatal("Penguin.Class() != aves")
	}
}

func TestPenguinConservation(t *testing.T) {
	if ok := Penguin.Conservation() == endangered; !ok {
		t.Fatal("Penguin.Conservation() != endangered")
	}
}

func TestPenguinDomain(t *testing.T) {
	if ok := Penguin.Domain() == eukarya; !ok {
		t.Fatal("Penguin.Domain() != eukarya")
	}
}

func TestPenguinFamily(t *testing.T) {
	if ok := Penguin.Family() == spheniscidae; !ok {
		t.Fatal("Penguin.Family() != spheniscidae")
	}
}

func TestPenguinGenus(t *testing.T) {
	if ok := Penguin.Genus() == na; !ok {
		t.Fatal("Penguin.Genus() != na")
	}
}

func TestPenguinKingdom(t *testing.T) {
	if ok := Penguin.Kingdom() == animalia; !ok {
		t.Fatal("Penguin.Kingdom() != animalia")
	}
}

func TestPenguinName(t *testing.T) {
	if ok := Penguin.Name() != na; !ok {
		t.Fatalf("Penguin.Name != %s", animals.Penguin.Name())
	}
}

func TestPenguinOrder(t *testing.T) {
	if ok := Penguin.Order() == sphenisciformes; !ok {
		t.Fatal("Penguin.Order() != sphenisciformes")
	}
}

func TestPenguinPhylum(t *testing.T) {
	if ok := Penguin.Phylum() == chordata; !ok {
		t.Fatal("Penguin.Phylum() != chordata")
	}
}

func TestPenguinSpecies(t *testing.T) {
	if ok := Penguin.Species() == na; !ok {
		t.Fatal("Penguin.Species() != na")
	}
}
