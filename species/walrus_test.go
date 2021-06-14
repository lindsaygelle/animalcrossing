package species

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestWalrusClass(t *testing.T) {
	if ok := Walrus.Class() == mammalia; !ok {
		t.Fatal("Walrus.Class() != mammalia")
	}
}

func TestWalrusConservation(t *testing.T) {
	if ok := Walrus.Conservation() == unknown; !ok {
		t.Fatal("Walrus.Conservation() != unknown")
	}
}

func TestWalrusDomain(t *testing.T) {
	if ok := Walrus.Domain() == eukarya; !ok {
		t.Fatal("Walrus.Domain() != eukarya")
	}
}

func TestWalrusFamily(t *testing.T) {
	if ok := Walrus.Family() == odobenidae; !ok {
		t.Fatal("Walrus.Family() != odobenidae")
	}
}

func TestWalrusGenus(t *testing.T) {
	if ok := Walrus.Genus() == odobenus; !ok {
		t.Fatal("Walrus.Genus() != odobenus")
	}
}

func TestWalrusKingdom(t *testing.T) {
	if ok := Walrus.Kingdom() == animalia; !ok {
		t.Fatal("Walrus.Kingdom() != animalia")
	}
}

func TestWalrusName(t *testing.T) {
	if ok := Walrus.Name() != na; !ok {
		t.Fatalf("Walrus.Name != %s", animals.Walrus.Name())
	}
}

func TestWalrusOrder(t *testing.T) {
	if ok := Walrus.Order() == carnivora; !ok {
		t.Fatal("Walrus.Order() != carnivora")
	}
}

func TestWalrusPhylum(t *testing.T) {
	if ok := Walrus.Phylum() == chordata; !ok {
		t.Fatal("Walrus.Phylum() != chordata")
	}
}

func TestWalrusSpecies(t *testing.T) {
	if ok := Walrus.Species() == oRosmarus; !ok {
		t.Fatal("Walrus.Species() != oRosmarus")
	}
}
