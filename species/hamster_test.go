package species

import (
	"github.com/lindsaygelle/animalcrossing/animals"
	"testing"
)

func TestHamsterClass(t *testing.T) {
	if ok := Hamster.Class() == mammalia; !ok {
		t.Fatal("Hamster.Class() != mammalia")
	}
}

func TestHamsterConservation(t *testing.T) {
	if ok := Hamster.Conservation() == leastConcern; !ok {
		t.Fatal("Hamster.Conservation() != leastConcern")
	}
}

func TestHamsterDomain(t *testing.T) {
	if ok := Hamster.Domain() == eukarya; !ok {
		t.Fatal("Hamster.Domain() != eukarya")
	}
}

func TestHamsterFamily(t *testing.T) {
	if ok := Hamster.Family() == cricetidae; !ok {
		t.Fatal("Hamster.Family() != cricetidae")
	}
}

func TestHamsterGenus(t *testing.T) {
	if ok := Hamster.Genus() == na; !ok {
		t.Fatal("Hamster.Genus() != na")
	}
}

func TestHamsterKingdom(t *testing.T) {
	if ok := Hamster.Kingdom() == animalia; !ok {
		t.Fatal("Hamster.Kingdom() != animalia")
	}
}

func TestHamsterName(t *testing.T) {
	if ok := Hamster.Name() != na; !ok {
		t.Fatalf("Hamster.Name != %s", animals.Hamster.Name())
	}
}

func TestHamsterOrder(t *testing.T) {
	if ok := Hamster.Order() == rodentia; !ok {
		t.Fatal("Hamster.Order() != rodentia")
	}
}

func TestHamsterPhylum(t *testing.T) {
	if ok := Hamster.Phylum() == chordata; !ok {
		t.Fatal("Hamster.Phylum() != chordata")
	}
}

func TestHamsterSpecies(t *testing.T) {
	if ok := Hamster.Species() == na; !ok {
		t.Fatal("Hamster.Species() != na")
	}
}
