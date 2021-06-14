package species

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestReindeerClass(t *testing.T) {
	if ok := Reindeer.Class() == mammalia; !ok {
		t.Fatal("Reindeer.Class() != mammalia")
	}
}

func TestReindeerConservation(t *testing.T) {
	if ok := Reindeer.Conservation() == vulnerable; !ok {
		t.Fatal("Reindeer.Conservation() != vulnerable")
	}
}

func TestReindeerDomain(t *testing.T) {
	if ok := Reindeer.Domain() == eukarya; !ok {
		t.Fatal("Reindeer.Domain() != eukarya")
	}
}

func TestReindeerFamily(t *testing.T) {
	if ok := Reindeer.Family() == cervidae; !ok {
		t.Fatal("Reindeer.Family() != cervidae")
	}
}

func TestReindeerGenus(t *testing.T) {
	if ok := Reindeer.Genus() == rangifer; !ok {
		t.Fatal("Reindeer.Genus() != rangifer")
	}
}

func TestReindeerKingdom(t *testing.T) {
	if ok := Reindeer.Kingdom() == animalia; !ok {
		t.Fatal("Reindeer.Kingdom() != animalia")
	}
}

func TestReindeerName(t *testing.T) {
	if ok := Reindeer.Name() != na; !ok {
		t.Fatalf("Reindeer.Name != %s", animals.Reindeer.Name())
	}
}

func TestReindeerOrder(t *testing.T) {
	if ok := Reindeer.Order() == artiodactyla; !ok {
		t.Fatal("Reindeer.Order() != artiodactyla")
	}
}

func TestReindeerPhylum(t *testing.T) {
	if ok := Reindeer.Phylum() == chordata; !ok {
		t.Fatal("Reindeer.Phylum() != chordata")
	}
}

func TestReindeerSpecies(t *testing.T) {
	if ok := Reindeer.Species() == rTarandus; !ok {
		t.Fatal("Reindeer.Species() != rTarandus")
	}
}
