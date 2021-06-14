package species

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestTortoiseClass(t *testing.T) {
	if ok := Tortoise.Class() == reptilia; !ok {
		t.Fatal("Tortoise.Class() != reptilia")
	}
}

func TestTortoiseConservation(t *testing.T) {
	if ok := Tortoise.Conservation() == unknown; !ok {
		t.Fatal("Tortoise.Conservation() != unknown")
	}
}

func TestTortoiseDomain(t *testing.T) {
	if ok := Tortoise.Domain() == eukarya; !ok {
		t.Fatal("Tortoise.Domain() != eukarya")
	}
}

func TestTortoiseFamily(t *testing.T) {
	if ok := Tortoise.Family() == testudinidae; !ok {
		t.Fatal("Tortoise.Family() != testudinidae")
	}
}

func TestTortoiseGenus(t *testing.T) {
	if ok := Tortoise.Genus() == na; !ok {
		t.Fatal("Tortoise.Genus() != na")
	}
}

func TestTortoiseKingdom(t *testing.T) {
	if ok := Tortoise.Kingdom() == animalia; !ok {
		t.Fatal("Tortoise.Kingdom() != animalia")
	}
}

func TestTortoiseName(t *testing.T) {
	if ok := Tortoise.Name() != na; !ok {
		t.Fatalf("Tortoise.Name != %s", animals.Tortoise.Name())
	}
}

func TestTortoiseOrder(t *testing.T) {
	if ok := Tortoise.Order() == testudines; !ok {
		t.Fatal("Tortoise.Order() != testudines")
	}
}

func TestTortoisePhylum(t *testing.T) {
	if ok := Tortoise.Phylum() == chordata; !ok {
		t.Fatal("Tortoise.Phylum() != chordata")
	}
}

func TestTortoiseSpecies(t *testing.T) {
	if ok := Tortoise.Species() == na; !ok {
		t.Fatal("Tortoise.Species() != na")
	}
}
