package species

import (
	"strings"
	"testing"
)

func TestFamilyAll(t *testing.T) {
	var (
		c = [...]string{
			accipitridae,
			agamidae,
			alligatoridae,
			ambystomatidae,
			anatidae,
			bovidae,
			camelidae,
			canidae,
			castoridae,
			cervidae,
			chamaeleonidae,
			cricetidae,
			columbidae,
			elephantidae,
			equidae,
			erinaceidae,
			felidae,
			giraffidae,
			hippopotamidea,
			hominidae,
			laridae,
			leporidae,
			macropodidae,
			mephitidae,
			muridae,
			mustelidae,
			odobenidae,
			otariidae,
			pelecanidae,
			phascolarctidae,
			phasianidae,
			rhinocerotidae,
			sciuridae,
			spheniscidae,
			struthionidae,
			suidae,
			talpidae,
			tapiridae,
			testudinidae,
			ursidae}
	)
	for i := 0; i < len(c); i++ {
		var s string = c[i]
		if ok := string(s[0]) == strings.ToUpper(string(s[0])); !ok {
			t.Fatalf("%s does not start with a capital letter", s)
		}
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

func TestFamilyCricetidae(t *testing.T) {
	var s string = "Cricetidae"
	if ok := cricetidae == s; !ok {
		t.Fatalf("cricetidae != %s", s)
	}
}

func TestFamilyElephantidae(t *testing.T) {
	var s string = "Elephantidae"
	if ok := elephantidae == s; !ok {
		t.Fatalf("elephantidae != %s", s)
	}
}

func TestFamilyEquidae(t *testing.T) {
	var s string = "Equidae"
	if ok := equidae == s; !ok {
		t.Fatalf("equidae != %s", s)
	}
}

func TestFamilyErinaceidae(t *testing.T) {
	var s string = "Erinaceidae"
	if ok := erinaceidae == s; !ok {
		t.Fatalf("erinaceidae != %s", s)
	}
}

func TestFamilyFelidae(t *testing.T) {
	var s string = "Felidae"
	if ok := felidae == s; !ok {
		t.Fatalf("felidae != %s", s)
	}
}

func TestFamilyGiraffidae(t *testing.T) {
	var s string = "Giraffidae"
	if ok := giraffidae == s; !ok {
		t.Fatalf("giraffidae != %s", s)
	}
}

func TestFamilyHippopotamidea(t *testing.T) {
	var s string = "Hippopotamidea"
	if ok := hippopotamidea == s; !ok {
		t.Fatalf("hippopotamidea != %s", s)
	}
}

func TestFamilyHominidae(t *testing.T) {
	var s string = "Hominidae"
	if ok := hominidae == s; !ok {
		t.Fatalf("hominidae != %s", s)
	}
}

func TestFamilyLaridae(t *testing.T) {
	var s string = "Laridae"
	if ok := laridae == s; !ok {
		t.Fatalf("laridae != %s", s)
	}
}

func TestFamilyLeporidae(t *testing.T) {
	var s string = "Leporidae"
	if ok := leporidae == s; !ok {
		t.Fatalf("leporidae != %s", s)
	}
}

func TestFamilyMacropodidae(t *testing.T) {
	var s string = "Macropodidae"
	if ok := macropodidae == s; !ok {
		t.Fatalf("macropodidae != %s", s)
	}
}

func TestFamilyMeleagris(t *testing.T) {
	var s string = "Meleagris"
	if ok := meleagris == s; !ok {
		t.Fatalf("meleagris != %s", s)
	}
}

func TestFamilyMephitidae(t *testing.T) {
	var s string = "Mephitidae"
	if ok := mephitidae == s; !ok {
		t.Fatalf("mephitidae != %s", s)
	}
}

func TestFamilyMuridae(t *testing.T) {
	var s string = "Muridae"
	if ok := muridae == s; !ok {
		t.Fatalf("muridae != %s", s)
	}
}

func TestFamilyMustelidae(t *testing.T) {
	var s string = "Mustelidae"
	if ok := mustelidae == s; !ok {
		t.Fatalf("mustelidae != %s", s)
	}
}

func TestFamilyOdobenidae(t *testing.T) {
	var s string = "Odobenidae"
	if ok := odobenidae == s; !ok {
		t.Fatalf("odobenidae != %s", s)
	}
}

func TestFamilyOtariidae(t *testing.T) {
	var s string = "Otariidae"
	if ok := otariidae == s; !ok {
		t.Fatalf("otariidae != %s", s)
	}
}

func TestFamilyPelecanidae(t *testing.T) {
	var s string = "Pelecanidae"
	if ok := pelecanidae == s; !ok {
		t.Fatalf("pelecanidae != %s", s)
	}
}

func TestFamilyPhascolarctidae(t *testing.T) {
	var s string = "Phascolarctidae"
	if ok := phascolarctidae == s; !ok {
		t.Fatalf("phascolarctidae != %s", s)
	}
}

func TestFamilyPhasianidae(t *testing.T) {
	var s string = "Phasianidae"
	if ok := phasianidae == s; !ok {
		t.Fatalf("phasianidae != %s", s)
	}
}

func TestFamilyRhinocerotidae(t *testing.T) {
	var s string = "Rhinocerotidae"
	if ok := rhinocerotidae == s; !ok {
		t.Fatalf("rhinocerotidae != %s", s)
	}
}

func TestFamilySciuridae(t *testing.T) {
	var s string = "Sciuridae"
	if ok := sciuridae == s; !ok {
		t.Fatalf("sciuridae != %s", s)
	}
}

func TestFamilySpheniscidae(t *testing.T) {
	var s string = "Spheniscidae"
	if ok := spheniscidae == s; !ok {
		t.Fatalf("spheniscidae != %s", s)
	}
}

func TestFamilyStruthionidae(t *testing.T) {
	var s string = "Struthionidae"
	if ok := struthionidae == s; !ok {
		t.Fatalf("struthionidae != %s", s)
	}
}

func TestFamilySuidae(t *testing.T) {
	var s string = "Suidae"
	if ok := suidae == s; !ok {
		t.Fatalf("suidae != %s", s)
	}
}

func TestFamilyTalpidae(t *testing.T) {
	var s string = "Talpidae"
	if ok := talpidae == s; !ok {
		t.Fatalf("talpidae != %s", talpidae)
	}
}

func TestFamilyTapiridae(t *testing.T) {
	var s string = "Tapiridae"
	if ok := tapiridae == s; !ok {
		t.Fatalf("tapiridae != %s", s)
	}
}

func TestFamilyTestudinidae(t *testing.T) {
	var s string = "Testudinidae"
	if ok := testudinidae == s; !ok {
		t.Fatalf("testudinidae != %s", s)
	}
}

func TestFamilyUrsidae(t *testing.T) {
	var s string = "Ursidae"
	if ok := ursidae == s; !ok {
		t.Fatalf("ursidae != %s", s)
	}
}
