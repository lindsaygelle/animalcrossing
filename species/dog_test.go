package species

import (
	"github.com/lindsaygelle/animalcrossing/animals"
	"testing"
)

func TestDogClass(t *testing.T) {
	if ok := Dog.Class() == mammalia; !ok {
		t.Fatal("Dog.Class() != mammalia")
	}
}

func TestDogConservation(t *testing.T) {
	if ok := Dog.Conservation() == domesticated; !ok {
		t.Fatal("Dog.Conservation() != domesticated")
	}
}

func TestDogDomain(t *testing.T) {
	if ok := Dog.Domain() == eukarya; !ok {
		t.Fatal("Dog.Domain() != eukarya")
	}
}

func TestDogFamily(t *testing.T) {
	if ok := Dog.Family() == canidae; !ok {
		t.Fatal("Dog.Family() != canidae")
	}
}

func TestDogGenus(t *testing.T) {
	if ok := Dog.Genus() == canis; !ok {
		t.Fatal("Dog.Genus() != canis")
	}
}

func TestDogKingdom(t *testing.T) {
	if ok := Dog.Kingdom() == animalia; !ok {
		t.Fatal("Dog.Kingdom() != animalia")
	}
}

func TestDogName(t *testing.T) {
	if ok := Dog.Name() != na; !ok {
		t.Fatalf("Dog.Name != %s", animals.Dog.Name())
	}
}

func TestDogOrder(t *testing.T) {
	if ok := Dog.Order() == carnivora; !ok {
		t.Fatal("Dog.Order() != carnivora")
	}
}

func TestDogPhylum(t *testing.T) {
	if ok := Dog.Phylum() == chordata; !ok {
		t.Fatal("Dog.Phylum() != chordata")
	}
}

func TestDogSpecies(t *testing.T) {
	if ok := Dog.Species() == canisLupusFamiliaris; !ok {
		t.Fatal("Dog.Species() != canisLupusFamiliaris")
	}
}
