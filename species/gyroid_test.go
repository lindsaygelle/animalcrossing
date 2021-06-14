package species

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestGyroidClass(t *testing.T) {
	if ok := Gyroid.Class() == na; !ok {
		t.Fatal("Gyroid.Class() != na")
	}
}

func TestGyroidConservation(t *testing.T) {
	if ok := Gyroid.Conservation() == unknown; !ok {
		t.Fatal("Gyroid.Conservation() != unknown")
	}
}

func TestGyroidDomain(t *testing.T) {
	if ok := Gyroid.Domain() == eukarya; !ok {
		t.Fatal("Gyroid.Domain() != eukarya")
	}
}

func TestGyroidFamily(t *testing.T) {
	if ok := Gyroid.Family() == na; !ok {
		t.Fatal("Gyroid.Family() != na")
	}
}

func TestGyroidFictional(t *testing.T) {
	if ok := Gyroid.Fictional() == fictional; !ok {
		t.Fatal("Gyroid.Fictional() != true")
	}
}

func TestGyroidGenus(t *testing.T) {
	if ok := Gyroid.Genus() == na; !ok {
		t.Fatal("Gyroid.Genus() != na")
	}
}

func TestGyroidKingdom(t *testing.T) {
	if ok := Gyroid.Kingdom() == na; !ok {
		t.Fatal("Gyroid.Kingdom() != na")
	}
}

func TestGyroidName(t *testing.T) {
	if ok := Gyroid.Name() != na; !ok {
		t.Fatalf("Gyroid.Name != %s", animals.Gyroid.Name())
	}
}

func TestGyroidOrder(t *testing.T) {
	if ok := Gyroid.Order() == na; !ok {
		t.Fatal("Gyroid.Order() != na")
	}
}

func TestGyroidPhylum(t *testing.T) {
	if ok := Gyroid.Phylum() == na; !ok {
		t.Fatal("Gyroid.Phylum() != na")
	}
}

func TestGyroidSpecies(t *testing.T) {
	if ok := Gyroid.Species() == na; !ok {
		t.Fatal("Gyroid.Species() != na")
	}
}
