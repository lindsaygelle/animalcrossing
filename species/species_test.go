package species

import "testing"

func TestClassAmphibia(t *testing.T) {
	var s string = "Amphibia"
	if ok := amphibia == s; !ok {
		t.Fatal("amphibia != Amphibia")
	}
}
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

func TestConservationCriticallyEndangered(t *testing.T) {
	var s string = "Critically Endangered"
	if ok := criticallyEndangered == s; !ok {
		t.Fatal("criticallyEndangered != Critically Endangered")
	}
}
func TestConservationDomesticated(t *testing.T) {
	var s string = "Domesticated"
	if ok := domesticated == s; !ok {
		t.Fatal("domesticated != Domesticated")
	}
}

func TestConservationVulnerable(t *testing.T) {
	var s string = "Vulnerable"
	if ok := vulnerable == s; !ok {
		t.Fatal("vulnerable != Vulnerable")
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

func TestFamilyAmbystomatidae(t *testing.T) {
	var s string = "Ambystomatidae"
	if ok := ambystomatidae == s; !ok {
		t.Fatal("ambystomatidae != Ambystomatidae")
	}
}

func TestGenusAlligator(t *testing.T) {
	var s string = "Alligator"
	if ok := alligator == s; !ok {
		t.Fatal("alligator != Alligator")
	}
}

func TestGenusAmbystoma(t *testing.T) {
	var s string = "Ambystoma"
	if ok := alligator == s; !ok {
		t.Fatal("ambystoma != Ambystoma")
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

func TestOrderCaudata(t *testing.T) {
	var s string = "Caudata"
	if ok := artiodactyla == s; !ok {
		t.Fatal("caudata != Caudata")
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

func TestSpeciesAMexicanum(t *testing.T) {
	var s string = "A. mexicanum"
	if ok := aMexicanum == s; !ok {
		t.Fatal("aMexicanum != A. mexicanum")
	}
}
