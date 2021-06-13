package species

import (
	"github.com/lindsaygelle/animalcrossing/animals"
	"testing"
)

func TestCatClass(t *testing.T) {
	if ok := Cat.Class() == mammalia; !ok {
		t.Fatal("Cat.Class() != mammalia")
	}
}

func TestCatConservation(t *testing.T) {
	if ok := Cat.Conservation() == domesticated; !ok {
		t.Fatal("Cat.Conservation() != domesticated")
	}
}

func TestCatDomain(t *testing.T) {
	if ok := Cat.Domain() == eukarya; !ok {
		t.Fatal("Cat.Domain() != eukarya")
	}
}

func TestCatFamily(t *testing.T) {
	if ok := Cat.Family() == felidae; !ok {
		t.Fatal("Cat.Family() != felidae")
	}
}

func TestCatGenus(t *testing.T) {
	if ok := Cat.Genus() == felis; !ok {
		t.Fatal("Cat.Genus() != felis")
	}
}

func TestCatKingdom(t *testing.T) {
	if ok := Cat.Kingdom() == animalia; !ok {
		t.Fatal("Cat.Kingdom() != animalia")
	}
}

func TestCatName(t *testing.T) {
	if ok := Cat.Name() != na; !ok {
		t.Fatalf("Cat.Name != %s", animals.Cat.Name())
	}
}

func TestCatOrder(t *testing.T) {
	if ok := Cat.Order() == carnivora; !ok {
		t.Fatal("Cat.Order() != carnivora")
	}
}

func TestCatPhylum(t *testing.T) {
	if ok := Cat.Phylum() == chordata; !ok {
		t.Fatal("Cat.Phylum() != chordata")
	}
}

func TestCatSpecies(t *testing.T) {
	if ok := Cat.Species() == fCatus; !ok {
		t.Fatal("Cat.Species() != fCatus")
	}
}
