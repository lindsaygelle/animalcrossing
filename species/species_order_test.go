package species

import (
	"strings"
	"testing"
)

func TestOrderAll(t *testing.T) {
	var (
		c = [...]string{accipitriformes,
			anseriformes,
			anura,
			artiodactyla,
			calumbiformes,
			carnivora,
			caudata,
			charadriiformes,
			columbiformes,
			crocodylia,
			diprotodontia,
			eulipotyphla,
			galliformes,
			lagomorpha,
			octopoda,
			pelecaniformes,
			perissodactyla,
			pilosa,
			primates,
			proboscidea,
			rodentia,
			sphenisciformes,
			squamata,
			strigiformes,
			struthioniformes,
			testudines}
	)
	for i := 0; i < len(c); i++ {
		var s string = c[i]
		if ok := string(s[0]) == strings.ToUpper(string(s[0])); !ok {
			t.Fatalf("%s does not start with a capital letter", s)
		}
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

func TestOrderCalumbiformes(t *testing.T) {
	var s string = "Calumbiformes"
	if ok := calumbiformes == s; !ok {
		t.Fatalf("calumbiformes != %s", s)
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

func TestOrderCharadriiformes(t *testing.T) {
	var s string = "Charadriiformes"
	if ok := charadriiformes == s; !ok {
		t.Fatalf("charadriiformes != %s", s)
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

func TestOrderDiprotodontia(t *testing.T) {
	var s string = "Diprotodontia"
	if ok := diprotodontia == s; !ok {
		t.Fatalf("diprotodontia != %s", s)
	}
}

func TestOrderEulipotyphla(t *testing.T) {
	var s string = "Eulipotyphla"
	if ok := eulipotyphla == s; !ok {
		t.Fatalf("eulipotyphla != %s", eulipotyphla)
	}
}

func TestOrderGalliformes(t *testing.T) {
	var s string = "Galliformes"
	if ok := galliformes == s; !ok {
		t.Fatalf("galliformes != %s", s)
	}
}

func TestOrderLagomorpha(t *testing.T) {
	var s string = "Lagomorpha"
	if ok := lagomorpha == s; !ok {
		t.Fatalf("lagomorpha != %s", s)
	}
}

func TestOrderOctopoda(t *testing.T) {
	var s string = "Octopoda"
	if ok := octopoda == s; !ok {
		t.Fatalf("octopoda != %s", s)
	}
}

func TestOrderPelecaniformes(t *testing.T) {
	var s string = "Pelecaniformes"
	if ok := pelecaniformes == s; !ok {
		t.Fatalf("pelecaniformes != %s", s)
	}
}

func TestOrderPerissodactyla(t *testing.T) {
	var s string = "Perissodactyla"
	if ok := perissodactyla == s; !ok {
		t.Fatalf("perissodactyla != %s", s)
	}
}

func TestOrderPilosa(t *testing.T) {
	var s string = "Pilosa"
	if ok := pilosa == s; !ok {
		t.Fatalf("pilosa != %s", s)
	}
}

func TestOrderPrimates(t *testing.T) {
	var s string = "Primates"
	if ok := primates == s; !ok {
		t.Fatalf("primates != %s", s)
	}
}

func TestOrderProboscidea(t *testing.T) {
	var s string = "Proboscidea"
	if ok := proboscidea == s; !ok {
		t.Fatalf("proboscidea != %s", s)
	}
}

func TestOrderSphenisciformes(t *testing.T) {
	var s string = "Sphenisciformes"
	if ok := sphenisciformes == s; !ok {
		t.Fatalf("sphenisciformes != %s", s)
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

func TestOrderStrigiformes(t *testing.T) {
	var s string = "Strigiformes"
	if ok := strigiformes == s; !ok {
		t.Fatalf("strigiformes != %s", s)
	}
}

func TestOrderStruthioniformes(t *testing.T) {
	var s string = "Struthioniformes"
	if ok := struthioniformes == s; !ok {
		t.Fatalf("struthioniformes != %s", s)
	}
}

func TestOrderTestudines(t *testing.T) {
	var s string = "Testudines"
	if ok := testudines == s; !ok {
		t.Fatalf("testudines != %s", s)
	}
}
