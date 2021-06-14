package species

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestRhinocerosClass(t *testing.T) {
	if ok := Rhinoceros.Class() == mammalia; !ok {
		t.Fatal("Rhinoceros.Class() != mammalia")
	}
}

func TestRhinocerosConservation(t *testing.T) {
	if ok := Rhinoceros.Conservation() == vulnerable; !ok {
		t.Fatal("Rhinoceros.Conservation() != vulnerable")
	}
}

func TestRhinocerosDomain(t *testing.T) {
	if ok := Rhinoceros.Domain() == eukarya; !ok {
		t.Fatal("Rhinoceros.Domain() != eukarya")
	}
}

func TestRhinocerosFamily(t *testing.T) {
	if ok := Rhinoceros.Family() == rhinocerotidae; !ok {
		t.Fatal("Rhinoceros.Family() != rhinocerotidae")
	}
}

func TestRhinocerosGenus(t *testing.T) {
	if ok := Rhinoceros.Genus() == na; !ok {
		t.Fatal("Rhinoceros.Genus() != na")
	}
}

func TestRhinocerosKingdom(t *testing.T) {
	if ok := Rhinoceros.Kingdom() == animalia; !ok {
		t.Fatal("Rhinoceros.Kingdom() != animalia")
	}
}

func TestRhinocerosName(t *testing.T) {
	if ok := Rhinoceros.Name() != na; !ok {
		t.Fatalf("Rhinoceros.Name != %s", animals.Rhinoceros.Name())
	}
}

func TestRhinocerosOrder(t *testing.T) {
	if ok := Rhinoceros.Order() == perissodactyla; !ok {
		t.Fatal("Rhinoceros.Order() != perissodactyla")
	}
}

func TestRhinocerosPhylum(t *testing.T) {
	if ok := Rhinoceros.Phylum() == chordata; !ok {
		t.Fatal("Rhinoceros.Phylum() != chordata")
	}
}

func TestRhinocerosSpecies(t *testing.T) {
	if ok := Rhinoceros.Species() == na; !ok {
		t.Fatal("Rhinoceros.Species() != na")
	}
}
