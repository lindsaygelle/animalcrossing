package species

import "testing"

func TestClassMammalia(t *testing.T) {
	var s string = "Mammalia"
	if ok := mammalia == s; !ok {
		t.Fatal("mammalia != Mammalia")
	}
}
func TestClassReptilia(t *testing.T) {
	var s string = "Reptilia"
	if ok := reptilia == s; !ok {
		t.Fatal("reptilia != Reptilia")
	}
}

func TestConservationVulnerable(t *testing.T) {
	var s string = "Vulnerable"
	if ok := vulnerable == s; !ok {
		t.Fatal("vulnerable != Vulnerable")
	}
}

func TestConservationDomesticated(t *testing.T) {
	var s string = "Domesticated"
	if ok := domesticated == s; !ok {
		t.Fatal("domesticated != Domesticated")
	}
}

func TestDomainEukarya(t *testing.T) {
	var s string = "Eukarya"
	if ok := eukarya == s; !ok {
		t.Fatal("eukarya != Eukarya")
	}
}

func TestFamilyAlligatoridae(t *testing.T) {
	var s string = "Alligatoridae"
	if ok := alligatoridae == s; !ok {
		t.Fatal("alligatoridae != Alligatoridae")
	}
}

func TestGenusAlligator(t *testing.T) {
	var s string = "Alligator"
	if ok := alligator == s; !ok {
		t.Fatal("alligator != Alligator")
	}
}

func TestGenusCamelidae(t *testing.T) {
	var s string = "Camelidae"
	if ok := camelidae == s; !ok {
		t.Fatal("camelidae != Camelidae")
	}
}

func TestKingdomAnimalia(t *testing.T) {
	var s string = "Animalia"
	if ok := animalia == s; !ok {
		t.Fatal("animalia != Animalia")
	}
}

func TestOrderArtiodactyla(t *testing.T) {
	var s string = "Artiodactyla"
	if ok := artiodactyla == s; !ok {
		t.Fatal("artiodactyla != Artiodactyla")
	}
}

func TestOrderPilosa(t *testing.T) {
	var s string = "Pilosa"
	if ok := pilosa == s; !ok {
		t.Fatal("pilosa != Pilosa")
	}
}

func TestOrderCrocodylia(t *testing.T) {
	var s string = "Crocodylia"
	if ok := crocodylia == s; !ok {
		t.Fatal("crocodylia != Crocodylia")
	}
}

func TestPhylumChordata(t *testing.T) {
	var s string = "Chordata"
	if ok := chordata == s; !ok {
		t.Fatal("chordata != Chordata")
	}
}
