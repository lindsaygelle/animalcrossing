package species

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestPumpkinClass(t *testing.T) {
	if ok := Pumpkin.Class() == na; !ok {
		t.Fatal("Pumpkin.Class() != na")
	}
}

func TestPumpkinConservation(t *testing.T) {
	if ok := Pumpkin.Conservation() == na; !ok {
		t.Fatal("Pumpkin.Conservation() != na")
	}
}

func TestPumpkinDomain(t *testing.T) {
	if ok := Pumpkin.Domain() == na; !ok {
		t.Fatal("Pumpkin.Domain() != na")
	}
}

func TestPumpkinFamily(t *testing.T) {
	if ok := Pumpkin.Family() == na; !ok {
		t.Fatal("Pumpkin.Family() != na")
	}
}

func TestPumpkinFictional(t *testing.T) {
	if ok := Pumpkin.Fictional() == (fictional); !ok {
		t.Fatal("Pumpkin.Fictional() != true")
	}
}

func TestPumpkinGenus(t *testing.T) {
	if ok := Pumpkin.Genus() == na; !ok {
		t.Fatal("Pumpkin.Genus() != na")
	}
}

func TestPumpkinKingdom(t *testing.T) {
	if ok := Pumpkin.Kingdom() == na; !ok {
		t.Fatal("Pumpkin.Kingdom() != na")
	}
}

func TestPumpkinName(t *testing.T) {
	if ok := Pumpkin.Name() == animals.Pumpkin.Name(); !ok {
		t.Fatalf("Pumpkin.Name != %s", animals.Pumpkin.Name())
	}
}

func TestPumpkinOrder(t *testing.T) {
	if ok := Pumpkin.Order() == na; !ok {
		t.Fatal("Pumpkin.Order() != na")
	}
}

func TestPumpkinPhylum(t *testing.T) {
	if ok := Pumpkin.Phylum() == na; !ok {
		t.Fatal("Pumpkin.Phylum() != na")
	}
}

func TestPumpkinSpecies(t *testing.T) {
	if ok := Pumpkin.Species() == na; !ok {
		t.Fatal("Pumpkin.Species() != na")
	}
}
