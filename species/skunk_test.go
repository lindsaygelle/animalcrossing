package species

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestSkunkClass(t *testing.T) {
	if ok := Skunk.Class() == mammalia; !ok {
		t.Fatal("Skunk.Class() != mammalia")
	}
}

func TestSkunkConservation(t *testing.T) {
	if ok := Skunk.Conservation() == unknown; !ok {
		t.Fatal("Skunk.Conservation() != unknown")
	}
}

func TestSkunkDomain(t *testing.T) {
	if ok := Skunk.Domain() == eukarya; !ok {
		t.Fatal("Skunk.Domain() != eukarya")
	}
}

func TestSkunkFamily(t *testing.T) {
	if ok := Skunk.Family() == mephitidae; !ok {
		t.Fatal("Skunk.Family() != mephitidae")
	}
}

func TestSkunkFictional(t *testing.T) {
	if ok := Skunk.Fictional() == (!fictional); !ok {
		t.Fatal("Skunk.Fictional() != false")
	}
}

func TestSkunkGenus(t *testing.T) {
	if ok := Skunk.Genus() == na; !ok {
		t.Fatal("Skunk.Genus() != na")
	}
}

func TestSkunkKingdom(t *testing.T) {
	if ok := Skunk.Kingdom() == animalia; !ok {
		t.Fatal("Skunk.Kingdom() != animalia")
	}
}

func TestSkunkName(t *testing.T) {
	if ok := Skunk.Name() != na; !ok {
		t.Fatalf("Skunk.Name != %s", animals.Skunk.Name())
	}
}

func TestSkunkOrder(t *testing.T) {
	if ok := Skunk.Order() == carnivora; !ok {
		t.Fatal("Skunk.Order() != carnivora")
	}
}

func TestSkunkPhylum(t *testing.T) {
	if ok := Skunk.Phylum() == chordata; !ok {
		t.Fatal("Skunk.Phylum() != chordata")
	}
}

func TestSkunkSpecies(t *testing.T) {
	if ok := Skunk.Species() == na; !ok {
		t.Fatal("Skunk.Species() != na")
	}
}
