package species

import (
	"github.com/lindsaygelle/animalcrossing/animals"
	"testing"
)

func TestChickenClass(t *testing.T) {
	if ok := Chicken.Class() == aves; !ok {
		t.Fatal("Chicken.Class() != aves")
	}
}

func TestChickenConservation(t *testing.T) {
	if ok := Chicken.Conservation() == domesticated; !ok {
		t.Fatal("Chicken.Conservation() != domesticated")
	}
}

func TestChickenDomain(t *testing.T) {
	if ok := Chicken.Domain() == eukarya; !ok {
		t.Fatal("Chicken.Domain() != eukarya")
	}
}

func TestChickenFamily(t *testing.T) {
	if ok := Chicken.Family() == phasianidae; !ok {
		t.Fatal("Chicken.Family() != phasianidae")
	}
}

func TestChickenGenus(t *testing.T) {
	if ok := Chicken.Genus() == gallus; !ok {
		t.Fatal("Chicken.Genus() != gallus")
	}
}

func TestChickenKingdom(t *testing.T) {
	if ok := Chicken.Kingdom() == animalia; !ok {
		t.Fatal("Chicken.Kingdom() != animalia")
	}
}

func TestChickenName(t *testing.T) {
	if ok := Chicken.Name() != na; !ok {
		t.Fatalf("Chicken.Name != %s", animals.Chicken.Name())
	}
}

func TestChickenOrder(t *testing.T) {
	if ok := Chicken.Order() == galliformes; !ok {
		t.Fatal("Chicken.Order() != galliformes")
	}
}

func TestChickenPhylum(t *testing.T) {
	if ok := Chicken.Phylum() == chordata; !ok {
		t.Fatal("Chicken.Phylum() != chordata")
	}
}

func TestChickenSpecies(t *testing.T) {
	if ok := Chicken.Species() == gGallus; !ok {
		t.Fatal("Chicken.Species() != gGallus")
	}
}
