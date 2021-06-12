package species

import "testing"

func TestAlpacaClass(t *testing.T) {
	if ok := Alpaca.Class() == mammalia; !ok {
		t.Fatal("Alpaca.Class() != mammalia")
	}
}
func TestAlpacaConservation(t *testing.T) {
	if ok := Alpaca.Conservation() == domesticated; !ok {
		t.Fatal("Alpaca.Conservation() != domesticated")
	}
}
func TestAlpacaDomain(t *testing.T) {
	if ok := Alpaca.Domain() == eukarya; !ok {
		t.Fatal("Alpaca.Domain() != eukarya")
	}
}
func TestAlpacaFamily(t *testing.T) {
	if ok := Alpaca.Family() == camelidae; !ok {
		t.Fatal("Alpaca.Family() != camelidae")
	}
}
func TestAlpacaGenus(t *testing.T) {
	if ok := Alpaca.Genus() == vicugna; !ok {
		t.Fatal("Alpaca.Genus() != vicugna")
	}
}
func TestAlpacaKingdom(t *testing.T) {
	if ok := Alpaca.Kingdom() == animalia; !ok {
		t.Fatal("Alpaca.Kingdom() != animalia")
	}
}

func TestAlpacaOrder(t *testing.T) {
	if ok := Alpaca.Order() == artiodactyla; !ok {
		t.Fatal("Alpaca.Order() != artiodactyla")
	}
}

func TestAlpacaPhylum(t *testing.T) {
	if ok := Alpaca.Phylum() == chordata; !ok {
		t.Fatal("Alpaca.Phylum() != chordata")
	}
}
