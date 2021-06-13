package species

import "testing"

func TestHorseClass(t *testing.T) {
	if ok := Horse.Class() == mammalia; !ok {
		t.Fatal("Horse.Class() != mammalia")
	}
}

func TestHorseConservation(t *testing.T) {
	if ok := Horse.Conservation() == domesticated; !ok {
		t.Fatal("Horse.Conservation() != domesticated")
	}
}
func TestHorseDomain(t *testing.T) {
	if ok := Horse.Domain() == eukarya; !ok {
		t.Fatal("Horse.Domain() != eukarya")
	}
}
func TestHorseFamily(t *testing.T) {
	if ok := Horse.Family() == equidae; !ok {
		t.Fatal("Horse.Family() != equidae")
	}
}
func TestHorseGenus(t *testing.T) {
	if ok := Horse.Genus() == equus; !ok {
		t.Fatal("Horse.Genus() != equus")
	}
}
func TestHorseKingdom(t *testing.T) {
	if ok := Horse.Kingdom() == animalia; !ok {
		t.Fatal("Horse.Kingdom() != animalia")
	}
}
func TestHorseOrder(t *testing.T) {
	if ok := Horse.Order() == perissodactyla; !ok {
		t.Fatal("Horse.Order() != perissodactyla")
	}
}
func TestHorsePhylum(t *testing.T) {
	if ok := Horse.Phylum() == chordata; !ok {
		t.Fatal("Horse.Phylum() != chordata")
	}
}
func TestHorseSpecies(t *testing.T) {
	if ok := Horse.Species() == ""; !ok {
		t.Fatal("Horse.Species() != \"\"")
	}
}
