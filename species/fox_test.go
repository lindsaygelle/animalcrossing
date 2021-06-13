package species

import (
	"github.com/lindsaygelle/animalcrossing/animals"
	"testing"
)

func TestFoxClass(t *testing.T) {
	if ok := Fox.Class() == mammalia; !ok {
		t.Fatal("Fox.Class() != mammalia")
	}
}

func TestFoxConservation(t *testing.T) {
	if ok := Fox.Conservation() == leastConcern; !ok {
		t.Fatal("Fox.Conservation() != leastConcern")
	}
}

func TestFoxDomain(t *testing.T) {
	if ok := Fox.Domain() == eukarya; !ok {
		t.Fatal("Fox.Domain() != eukarya")
	}
}

func TestFoxFamily(t *testing.T) {
	if ok := Fox.Family() == canidae; !ok {
		t.Fatal("Fox.Family() != canidae")
	}
}

func TestFoxGenus(t *testing.T) {
	if ok := Fox.Genus() == vulpes; !ok {
		t.Fatal("Fox.Genus() != vulpes")
	}
}

func TestFoxName(t *testing.T) {
	if ok := Fox.Name() != na; !ok {
		t.Fatalf("Fox.Name != %s", animals.Fox.Name())
	}
}

func TestFoxKingdom(t *testing.T) {
	if ok := Fox.Kingdom() == animalia; !ok {
		t.Fatal("Fox.Kingdom() != animalia")
	}
}

func TestFoxOrder(t *testing.T) {
	if ok := Fox.Order() == carnivora; !ok {
		t.Fatal("Fox.Order() != carnivora")
	}
}

func TestFoxPhylum(t *testing.T) {
	if ok := Fox.Phylum() == chordata; !ok {
		t.Fatal("Fox.Phylum() != chordata")
	}
}

func TestFoxSpecies(t *testing.T) {
	if ok := Fox.Species() == vVulpes; !ok {
		t.Fatal("Fox.Species() != vVulpes")
	}
}
