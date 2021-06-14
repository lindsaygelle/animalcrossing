package species

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestFrillNeckedLizardClass(t *testing.T) {
	if ok := FrillNeckedLizard.Class() == sauropsida; !ok {
		t.Fatal("FrillNeckedLizard.Class() != sauropsida")
	}
}

func TestFrillNeckedLizardConservation(t *testing.T) {
	if ok := FrillNeckedLizard.Conservation() == leastConcern; !ok {
		t.Fatal("FrillNeckedLizard.Conservation() != leastConcern")
	}
}

func TestFrillNeckedLizardDomain(t *testing.T) {
	if ok := FrillNeckedLizard.Domain() == eukarya; !ok {
		t.Fatal("FrillNeckedLizard.Domain() != eukarya")
	}
}

func TestFrillNeckedLizardFamily(t *testing.T) {
	if ok := FrillNeckedLizard.Family() == agamidae; !ok {
		t.Fatal("FrillNeckedLizard.Family() != agamidae")
	}
}

func TestFrillNeckedLizardFictional(t *testing.T) {
	if ok := FrillNeckedLizard.Fictional() == (!fictional); !ok {
		t.Fatal("FrillNeckedLizard.Fictional() != false")
	}
}

func TestFrillNeckedLizardGenus(t *testing.T) {
	if ok := FrillNeckedLizard.Genus() == chlamydosaurus; !ok {
		t.Fatal("FrillNeckedLizard.Genus() != chlamydosaurus")
	}
}

func TestFrillNeckedLizardKingdom(t *testing.T) {
	if ok := FrillNeckedLizard.Kingdom() == animalia; !ok {
		t.Fatal("FrillNeckedLizard.Kingdom() != animalia")
	}
}

func TestFrillNeckedLizardName(t *testing.T) {
	if ok := FrillNeckedLizard.Name() != na; !ok {
		t.Fatalf("FrillNeckedLizard.Name != %s", animals.FrillNeckedLizard.Name())
	}
}

func TestFrillNeckedLizardOrder(t *testing.T) {
	if ok := FrillNeckedLizard.Order() == squamata; !ok {
		t.Fatal("FrillNeckedLizard.Order() != squamata")
	}
}

func TestFrillNeckedLizardPhylum(t *testing.T) {
	if ok := FrillNeckedLizard.Phylum() == chordata; !ok {
		t.Fatal("FrillNeckedLizard.Phylum() != chordata")
	}
}

func TestFrillNeckedLizardSpecies(t *testing.T) {
	if ok := FrillNeckedLizard.Species() == cKingii; !ok {
		t.Fatal("FrillNeckedLizard.Species() != cKingii")
	}
}
