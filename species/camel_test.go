package species

import (
	"github.com/lindsaygelle/animalcrossing/animals"
	"testing"
)

func TestCamelClass(t *testing.T) {
	if ok := Camel.Class() == mammalia; !ok {
		t.Fatal("Camel.Class() != mammalia")
	}
}

func TestCamelConservation(t *testing.T) {
	if ok := Camel.Conservation() == domesticated; !ok {
		t.Fatal("Camel.Conservation() != domesticated")
	}
}

func TestCamelDomain(t *testing.T) {
	if ok := Camel.Domain() == eukarya; !ok {
		t.Fatal("Camel.Domain() != eukarya")
	}
}

func TestCamelFamily(t *testing.T) {
	if ok := Camel.Family() == camelidae; !ok {
		t.Fatal("Camel.Family() != camelidae")
	}
}

func TestCamelGenus(t *testing.T) {
	if ok := Camel.Genus() == camelus; !ok {
		t.Fatal("Camel.Genus() != camelus")
	}
}

func TestCamelKingdom(t *testing.T) {
	if ok := Camel.Kingdom() == animalia; !ok {
		t.Fatal("Camel.Kingdom() != animalia")
	}
}

func TestCamelName(t *testing.T) {
	if ok := Camel.Name() != na; !ok {
		t.Fatalf("Camel.Name != %s", animals.Camel.Name())
	}
}

func TestCamelOrder(t *testing.T) {
	if ok := Camel.Order() == artiodactyla; !ok {
		t.Fatal("Camel.Order() != artiodactyla")
	}
}

func TestCamelPhylum(t *testing.T) {
	if ok := Camel.Phylum() == chordata; !ok {
		t.Fatal("Camel.Phylum() != chordata")
	}
}

func TestCamelSpecies(t *testing.T) {
	if ok := Camel.Species() == na; !ok {
		t.Fatal("Camel.Species() != na")
	}
}
