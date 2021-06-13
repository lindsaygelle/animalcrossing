package species

import "testing"

func TestClassAmphibia(t *testing.T) {
	var s string = "Amphibia"
	if ok := amphibia == s; !ok {
		t.Fatalf("amphibia != %s", s)
	}
}

func TestClassAves(t *testing.T) {
	var s string = "Aves"
	if ok := aves == s; !ok {
		t.Fatalf("aves != %s", s)
	}
}

func TestClassMammalia(t *testing.T) {
	var s string = "Mammalia"
	if ok := mammalia == s; !ok {
		t.Fatalf("mammalia != %s", s)
	}
}
func TestClassReptilia(t *testing.T) {
	var s string = "Reptilia"
	if ok := reptilia == s; !ok {
		t.Fatalf("reptilia != %s", s)
	}
}

func TestClassSauropsida(t *testing.T) {
	var s string = "Sauropsida"
	if ok := sauropsida == s; !ok {
		t.Fatalf("sauropsida != %s", s)
	}
}

func TestConservationCriticallyEndangered(t *testing.T) {
	var s string = "Critically Endangered"
	if ok := criticallyEndangered == s; !ok {
		t.Fatalf("criticallyEndangered != %s", s)
	}
}
func TestConservationDomesticated(t *testing.T) {
	var s string = "Domesticated"
	if ok := domesticated == s; !ok {
		t.Fatalf("domesticated != %s", s)
	}
}

func TestConservationEndangered(t *testing.T) {
	var s string = "Endangered"
	if ok := endangered == s; !ok {
		t.Fatalf("endangered != %s", s)
	}
}

func TestConservationExtinct(t *testing.T) {
	var s string = "Extinct"
	if ok := extinct == s; !ok {
		t.Fatalf("extinct != %s", s)
	}
}

func TestConservationLeastConcern(t *testing.T) {
	var s string = "Least Concern"
	if ok := leastConcern == s; !ok {
		t.Fatalf("leastConcern != %s", s)
	}
}

func TestConservationVulnerable(t *testing.T) {
	var s string = "Vulnerable"
	if ok := vulnerable == s; !ok {
		t.Fatalf("vulnerable != %s", s)
	}
}
func TestDomainEukarya(t *testing.T) {
	var s string = "Eukarya"
	if ok := eukarya == s; !ok {
		t.Fatalf("eukarya != %s", s)
	}
}

func TestFamilyAccipitridae(t *testing.T) {
	var s string = "Accipitridae"
	if ok := accipitridae == s; !ok {
		t.Fatalf("accipitridae != %s", s)
	}
}

func TestFamilyAgamidae(t *testing.T) {
	var s string = "Agamidae"
	if ok := agamidae == s; !ok {
		t.Fatalf("agamidae != %s", s)
	}
}

func TestFamilyAlligatoridae(t *testing.T) {
	var s string = "Alligatoridae"
	if ok := alligatoridae == s; !ok {
		t.Fatalf("alligatoridae != %s", s)
	}
}

func TestFamilyAmbystomatidae(t *testing.T) {
	var s string = "Ambystomatidae"
	if ok := ambystomatidae == s; !ok {
		t.Fatalf("ambystomatidae != %s", s)
	}
}

func TestFamilyAnatidae(t *testing.T) {
	var s string = "Anatidae"
	if ok := anatidae == s; !ok {
		t.Fatalf("anatidae != %s", s)
	}
}

func TestFamilyBovidae(t *testing.T) {
	var s string = "Bovidae"
	if ok := bovidae == s; !ok {
		t.Fatalf("bovidae != %s", s)
	}
}

func TestFamilyCamelidae(t *testing.T) {
	var s string = "Camelidae"
	if ok := camelidae == s; !ok {
		t.Fatalf("camelidae != %s", s)
	}
}

func TestFamilyCanidae(t *testing.T) {
	var s string = "Canidae"
	if ok := canidae == s; !ok {
		t.Fatalf("canidae != %s", s)
	}
}

func TestFamilyCastoridae(t *testing.T) {
	var s string = "Castoridae"
	if ok := castoridae == s; !ok {
		t.Fatalf("castoridae != %s", s)
	}
}

func TestFamilyCervidae(t *testing.T) {
	var s string = "Cervidae"
	if ok := cervidae == s; !ok {
		t.Fatalf("cervidae != %s", s)
	}
}

func TestFamilyChamaeleonidae(t *testing.T) {
	var s string = "Chamaeleonidae"
	if ok := chamaeleonidae == s; !ok {
		t.Fatalf("chamaeleonidae != %s", s)
	}
}

func TestFamilyColumbidae(t *testing.T) {
	var s string = "Columbidae"
	if ok := columbidae == s; !ok {
		t.Fatalf("columbidae != %s", s)
	}
}

func TestFamilyElephantidae(t *testing.T) {
	var s string = "Elephantidae"
	if ok := elephantidae == s; !ok {
		t.Fatalf("elephantidae != %s", s)
	}
}

func TestFamilyFelidae(t *testing.T) {
	var s string = "Felidae"
	if ok := felidae == s; !ok {
		t.Fatalf("felidae != %s", s)
	}
}

func TestFamilyPhasianidae(t *testing.T) {
	var s string = "Phasianidae"
	if ok := phasianidae == s; !ok {
		t.Fatalf("phasianidae != %s", s)
	}
}

func TestFamilySuidae(t *testing.T) {
	var s string = "Suidae"
	if ok := suidae == s; !ok {
		t.Fatalf("suidae != %s", s)
	}
}

func TestFamilyUrsidae(t *testing.T) {
	var s string = "Ursidae"
	if ok := ursidae == s; !ok {
		t.Fatalf("ursidae != %s", s)
	}
}

func TestGenusAlligator(t *testing.T) {
	var s string = "Alligator"
	if ok := alligator == s; !ok {
		t.Fatalf("alligator != %s", s)
	}
}

func TestGenusAmbystoma(t *testing.T) {
	var s string = "Ambystoma"
	if ok := ambystoma == s; !ok {
		t.Fatalf("ambystoma != %s", s)
	}
}

func TestGenusBos(t *testing.T) {
	var s string = "Bos"
	if ok := bos == s; !ok {
		t.Fatalf("bos != %s", s)
	}
}

func TestGenusCamelus(t *testing.T) {
	var s string = "Camelus"
	if ok := camelus == s; !ok {
		t.Fatalf("camelus != %s", s)
	}
}

func TestGenusCanis(t *testing.T) {
	var s string = "Canis"
	if ok := canis == s; !ok {
		t.Fatalf("canis != %s", s)
	}
}

func TestGenusCastor(t *testing.T) {
	var s string = "Castor"
	if ok := castor == s; !ok {
		t.Fatalf("castor != %s", s)
	}
}

func TestGenusChlamydosaurus(t *testing.T) {
	var s string = "Chlamydosaurus"
	if ok := chlamydosaurus == s; !ok {
		t.Fatalf("chlamydosaurus != %s", s)
	}
}

func TestGenusFelis(t *testing.T) {
	var s string = "Felis"
	if ok := felis == s; !ok {
		t.Fatalf("felis != %s", s)
	}
}

func TestGenusGallus(t *testing.T) {
	var s string = "Gallus"
	if ok := gallus == s; !ok {
		t.Fatalf("gallus != %s", s)
	}
}

func TestGenusRaphus(t *testing.T) {
	var s string = "Raphus"
	if ok := raphus == s; !ok {
		t.Fatalf("raphus != %s", s)
	}
}

func TestGenusSus(t *testing.T) {
	var s string = "Sus"
	if ok := sus == s; !ok {
		t.Fatalf("sus != %s", s)
	}
}

func TestGenusVicugna(t *testing.T) {
	var s string = "Vicugna"
	if ok := vicugna == s; !ok {
		t.Fatalf("vicugna != %s", s)
	}
}

func TestGenusVulpes(t *testing.T) {
	var s string = "Vulpes"
	if ok := vulpes == s; !ok {
		t.Fatalf("vulpes != %s", s)
	}
}

func TestKingdomAnimalia(t *testing.T) {
	var s string = "Animalia"
	if ok := animalia == s; !ok {
		t.Fatalf("animalia != %s", s)
	}
}

func TestKingdomEnimalia(t *testing.T) {
	var s string = "Enimalia"
	if ok := enimalia == s; !ok {
		t.Fatalf("enimalia != %s", s)
	}
}

func TestOrderAccipitriformes(t *testing.T) {
	var s string = "Accipitriformes"
	if ok := accipitriformes == s; !ok {
		t.Fatalf("accipitriformes != %s", s)
	}
}

func TestOrderAnseriformes(t *testing.T) {
	var s string = "Anseriformes"
	if ok := anseriformes == s; !ok {
		t.Fatalf("anseriformes != %s", s)
	}
}

func TestOrderAnura(t *testing.T) {
	var s string = "Anura"
	if ok := anura == s; !ok {
		t.Fatalf("anura != %s", s)
	}
}

func TestOrderArtiodactyla(t *testing.T) {
	var s string = "Artiodactyla"
	if ok := artiodactyla == s; !ok {
		t.Fatalf("artiodactyla != %s", s)
	}
}

func TestOrderCarnivora(t *testing.T) {
	var s string = "Carnivora"
	if ok := carnivora == s; !ok {
		t.Fatalf("carnivora != %s", s)
	}
}

func TestOrderCaudata(t *testing.T) {
	var s string = "Caudata"
	if ok := caudata == s; !ok {
		t.Fatalf("caudata != %s", s)
	}
}

func TestOrderColumbiformes(t *testing.T) {
	var s string = "Columbiformes"
	if ok := columbiformes == s; !ok {
		t.Fatalf("columbiformes != %s", s)
	}
}

func TestOrderCrocodylia(t *testing.T) {
	var s string = "Crocodylia"
	if ok := crocodylia == s; !ok {
		t.Fatalf("crocodylia != %s", s)
	}
}

func TestOrderGalliformes(t *testing.T) {
	var s string = "Galliformes"
	if ok := galliformes == s; !ok {
		t.Fatalf("galliformes != %s", s)
	}
}

func TestOrderPilosa(t *testing.T) {
	var s string = "Pilosa"
	if ok := pilosa == s; !ok {
		t.Fatalf("pilosa != %s", s)
	}
}

func TestOrderProboscidea(t *testing.T) {
	var s string = "Proboscidea"
	if ok := proboscidea == s; !ok {
		t.Fatalf("proboscidea != %s", s)
	}
}

func TestOrderRodentia(t *testing.T) {
	var s string = "Rodentia"
	if ok := rodentia == s; !ok {
		t.Fatalf("rodentia != %s", s)
	}
}

func TestOrderSquamata(t *testing.T) {
	var s string = "Squamata"
	if ok := squamata == s; !ok {
		t.Fatalf("squamata != %s", s)
	}
}

func TestPhylumChordata(t *testing.T) {
	var s string = "Chordata"
	if ok := chordata == s; !ok {
		t.Fatalf("chordata != %s", s)
	}
}

func TestSpeciesAMexicanum(t *testing.T) {
	var s string = "A. mexicanum"
	if ok := aMexicanum == s; !ok {
		t.Fatalf("aMexicanum != %s", s)
	}
}

func TestSpeciesBTaurus(t *testing.T) {
	var s string = "B. taurus"
	if ok := bTaurus == s; !ok {
		t.Fatalf("bTaurus != %s", s)
	}
}

func TestSpeciesCanisLupusFamiliaris(t *testing.T) {
	var s string = "Canis lupus familiaris"
	if ok := canisLupusFamiliaris == s; !ok {
		t.Fatalf("canisLupusFamiliaris != %s", s)
	}
}

func TestSpeciesCKingii(t *testing.T) {
	var s string = "C. kingii"
	if ok := cKingii == s; !ok {
		t.Fatalf("cKingii != %s", s)
	}
}

func TestSpeciesFCatus(t *testing.T) {
	var s string = "F. catus"
	if ok := fCatus == s; !ok {
		t.Fatalf("fCatus != %s", s)
	}
}

func TestSpeciesGGallus(t *testing.T) {
	var s string = "G. gallus"
	if ok := gGallus == s; !ok {
		t.Fatalf("gGallus != %s", s)
	}
}

func TestSpeciesRCucullatus(t *testing.T) {
	var s string = "R. cucullatus"
	if ok := rCucullatus == s; !ok {
		t.Fatalf("rCucullatus != %s", s)
	}
}

func TestSpeciesSScrofa(t *testing.T) {
	var s string = "S. scrofa"
	if ok := sScrofa == s; !ok {
		t.Fatalf("sScrofa != %s", s)
	}
}

func TestSpeciesVVulpes(t *testing.T) {
	var s string = "V. vulpes"
	if ok := vVulpes == s; !ok {
		t.Fatalf("vVulpes != %s", s)
	}
}
