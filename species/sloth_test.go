package species

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestSlothClass(t *testing.T) {
	if ok := Sloth.Class() == mammalia; !ok {
		t.Fatal("Sloth.Class() != mammalia")
	}
}

func TestSlothConservation(t *testing.T) {
	if ok := Sloth.Conservation() == unknown; !ok {
		t.Fatal("Sloth.Conservation() != unknown")
	}
}

func TestSlothDomain(t *testing.T) {
	if ok := Sloth.Domain() == eukarya; !ok {
		t.Fatal("Sloth.Domain() != eukarya")
	}
}

func TestSlothFamily(t *testing.T) {
	if ok := Sloth.Family() == na; !ok {
		t.Fatal("Sloth.Family() != na")
	}
}

func TestSlothFictional(t *testing.T) {
	if ok := Sloth.Fictional() == (!fictional); !ok {
		t.Fatal("Sloth.Fictional() != false")
	}
}

func TestSlothGenus(t *testing.T) {
	if ok := Sloth.Genus() == na; !ok {
		t.Fatal("Sloth.Genus() != na")
	}
}

func TestSlothKingdom(t *testing.T) {
	if ok := Sloth.Kingdom() == animalia; !ok {
		t.Fatal("Sloth.Kingdom() != animalia")
	}
}

func TestSlothName(t *testing.T) {
	if ok := Sloth.Name() != na; !ok {
		t.Fatalf("Sloth.Name != %s", animals.Sloth.Name())
	}
}

func TestSlothOrder(t *testing.T) {
	if ok := Sloth.Order() == pilosa; !ok {
		t.Fatal("Sloth.Order() != pilosa")
	}
}

func TestSlothPhylum(t *testing.T) {
	if ok := Sloth.Phylum() == chordata; !ok {
		t.Fatal("Sloth.Phylum() != chordata")
	}
}

func TestSlothSpecies(t *testing.T) {
	if ok := Sloth.Species() == na; !ok {
		t.Fatal("Sloth.Species() != na")
	}
}
