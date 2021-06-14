package species

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestPigClass(t *testing.T) {
	if ok := Pig.Class() == mammalia; !ok {
		t.Fatal("Pig.Class() != mammalia")
	}
}

func TestPigConservation(t *testing.T) {
	if ok := Pig.Conservation() == domesticated; !ok {
		t.Fatal("Pig.Conservation() != domesticated")
	}
}

func TestPigDomain(t *testing.T) {
	if ok := Pig.Domain() == eukarya; !ok {
		t.Fatal("Pig.Domain() != eukarya")
	}
}

func TestPigFamily(t *testing.T) {
	if ok := Pig.Family() == suidae; !ok {
		t.Fatal("Pig.Family() != suidae")
	}
}

func TestPigFictional(t *testing.T) {
	if ok := Pig.Fictional() == (!fictional); !ok {
		t.Fatal("Pig.Fictional() != false")
	}
}

func TestPigGenus(t *testing.T) {
	if ok := Pig.Genus() == sus; !ok {
		t.Fatal("Pig.Genus() != sus")
	}
}

func TestPigKingdom(t *testing.T) {
	if ok := Pig.Kingdom() == animalia; !ok {
		t.Fatal("Pig.Kingdom() != animalia")
	}
}

func TestPigName(t *testing.T) {
	if ok := Pig.Name() != na; !ok {
		t.Fatalf("Pig.Name != %s", animals.Pig.Name())
	}
}

func TestPigOrder(t *testing.T) {
	if ok := Pig.Order() == artiodactyla; !ok {
		t.Fatal("Pig.Order() != artiodactyla")
	}
}

func TestPigPhylum(t *testing.T) {
	if ok := Pig.Phylum() == chordata; !ok {
		t.Fatal("Pig.Phylum() != chordata")
	}
}

func TestPigSpecies(t *testing.T) {
	if ok := Pig.Species() == na; !ok {
		t.Fatal("Pig.Species() != na")
	}
}
