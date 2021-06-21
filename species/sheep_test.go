package species

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestSheepClass(t *testing.T) {
	if ok := Sheep.Class() == mammalia; !ok {
		t.Fatal("Sheep.Class() != mammalia")
	}
}

func TestSheepConservation(t *testing.T) {
	if ok := Sheep.Conservation() == domesticated; !ok {
		t.Fatal("Sheep.Conservation() != domesticated")
	}
}

func TestSheepDomain(t *testing.T) {
	if ok := Sheep.Domain() == eukarya; !ok {
		t.Fatal("Sheep.Domain() != eukarya")
	}
}

func TestSheepFamily(t *testing.T) {
	if ok := Sheep.Family() == bovidae; !ok {
		t.Fatal("Sheep.Family() != bovidae")
	}
}

func TestSheepFictional(t *testing.T) {
	if ok := Sheep.Fictional() == (!fictional); !ok {
		t.Fatal("Sheep.Fictional() != false")
	}
}

func TestSheepGenus(t *testing.T) {
	if ok := Sheep.Genus() == ovis; !ok {
		t.Fatal("Sheep.Genus() != ovis")
	}
}

func TestSheepKingdom(t *testing.T) {
	if ok := Sheep.Kingdom() == animalia; !ok {
		t.Fatal("Sheep.Kingdom() != animalia")
	}
}

func TestSheepName(t *testing.T) {
	if ok := Sheep.Name() == animals.Sheep.Name(); !ok {
		t.Fatalf("Sheep.Name != %s", animals.Sheep.Name())
	}
}

func TestSheepOrder(t *testing.T) {
	if ok := Sheep.Order() == artiodactyla; !ok {
		t.Fatal("Sheep.Order() != artiodactyla")
	}
}

func TestSheepPhylum(t *testing.T) {
	if ok := Sheep.Phylum() == chordata; !ok {
		t.Fatal("Sheep.Phylum() != chordata")
	}
}

func TestSheepSpecies(t *testing.T) {
	if ok := Sheep.Species() == oAries; !ok {
		t.Fatal("Sheep.Species() != oAries")
	}
}
