package species

import "testing"

func TestClassAmphibia(t *testing.T) {
	var s string = "Amphibia"
	if ok := amphibia == s; !ok {
		t.Fatal("amphibia != Amphibia")
	}
}

func TestClassAves(t *testing.T) {
	var s string = "Aves"
	if ok := aves == s; !ok {
		t.Fatal("aves != Aves")
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

func TestConservationLeastConcern(t *testing.T) {
	var s string = "Least Concern"
	if ok := leastConcern == s; !ok {
		t.Fatal("leastConcern != Least Concern")
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

func TestFamilyBovidae(t *testing.T) {
	var s string = "Bovidae"
	if ok := bovidae == s; !ok {
		t.Fatal("bovidae != Bovidae")
	}
}

func TestFamilyCastoridae(t *testing.T) {
	var s string = "Castoridae"
	if ok := castoridae == s; !ok {
		t.Fatal("castoridae != Castoridae")
	}
}

func TestFamilySuidae(t *testing.T) {
	var s string = "Suidae"
	if ok := suidae == s; !ok {
		t.Fatal("suidae != Suidae")
	}
}

func TestFamilyUrsidae(t *testing.T) {
	var s string = "Ursidae"
	if ok := ursidae == s; !ok {
		t.Fatal("ursidae != Ursidae")
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
	if ok := ambystoma == s; !ok {
		t.Fatal("ambystoma != Ambystoma")
	}
}

func TestGenusBos(t *testing.T) {
	var s string = "Bos"
	if ok := bos == s; !ok {
		t.Fatal("bos != Bos")
	}
}

func TestGenusCamelidae(t *testing.T) {
	var s string = "Camelidae"
	if ok := camelidae == s; !ok {
		t.Fatal("camelidae != Camelidae")
	}
}

func TestGenusCastor(t *testing.T) {
	var s string = "Castor"
	if ok := castor == s; !ok {
		t.Fatal("castor != Castor")
	}
}

func TestGenusSus(t *testing.T) {
	var s string = "Sus"
	if ok := sus == s; !ok {
		t.Fatal("sus != Sus")
	}
}

func TestGenusVicugna(t *testing.T) {
	var s string = "Vicugna"
	if ok := vicugna == s; !ok {
		t.Fatal("vicugna != Vicugna")
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

func TestOrderCarnivora(t *testing.T) {
	var s string = "Carnivora"
	if ok := carnivora == s; !ok {
		t.Fatal("carnivora != Carnivora")
	}
}

func TestOrderCaudata(t *testing.T) {
	var s string = "Caudata"
	if ok := caudata == s; !ok {
		t.Fatal("caudata != Caudata")
	}
}

func TestOrderCrocodylia(t *testing.T) {
	var s string = "Crocodylia"
	if ok := crocodylia == s; !ok {
		t.Fatal("crocodylia != Crocodylia")
	}
}

func TestOrderPilosa(t *testing.T) {
	var s string = "Pilosa"
	if ok := pilosa == s; !ok {
		t.Fatal("pilosa != Pilosa")
	}
}

func TestOrderRodentia(t *testing.T) {
	var s string = "Rodentia"
	if ok := rodentia == s; !ok {
		t.Fatal("rodentia != Rodentia")
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

func TestSpeciesBTaurus(t *testing.T) {
	var s string = "B. taurus"
	if ok := bTaurus == s; !ok {
		t.Fatal("bTaurus != B. taurus")
	}
}

func TestSpeicesSScrofa(t *testing.T) {
	var s string = "S. scrofa"
	if ok := sScrofa == s; !ok {
		t.Fatal("sScrofa != S. scrofa")
	}
}
