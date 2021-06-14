package species

import (
	"strings"
	"testing"
)

func TestGenusAll(t *testing.T) {
	var (
		c = [...]string{alligator,
			ambystoma,
			bos,
			camelus,
			canis,
			capra,
			castor,
			chlamydosaurus,
			equus,
			felis,
			gallus,
			giraffa,
			gorilla,
			hippopotamus,
			macropus,
			meleagris,
			mus,
			nyctereutes,
			odobenus,
			ovis,
			panthera,
			pavo,
			pelecanus,
			phascolarctos,
			raphus,
			rangifer,
			struthio,
			sus,
			tapirus,
			vicugna,
			vulpes}
	)
	for i := 0; i < len(c); i++ {
		var s string = c[i]
		if ok := string(s[0]) == strings.ToUpper(string(s[0])); !ok {
			t.Fatalf("%s does not start with a capital letter", s)
		}
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

func TestGenusCapra(t *testing.T) {
	var s string = "Capra"
	if ok := capra == s; !ok {
		t.Fatalf("capra != %s", s)
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

func TestGenusEquus(t *testing.T) {
	var s string = "Equus"
	if ok := equus == s; !ok {
		t.Fatalf("equus != %s", s)
	}
}

func TestGenusFelis(t *testing.T) {
	var s string = "Felis"
	if ok := felis == s; !ok {
		t.Fatalf("felis != %s", s)
	}
}

func TestGenusHippopotamus(t *testing.T) {
	var s string = "Hippopotamus"
	if ok := hippopotamus == s; !ok {
		t.Fatalf("hippopotamus != %s", s)
	}
}

func TestGenusGallus(t *testing.T) {
	var s string = "Gallus"
	if ok := gallus == s; !ok {
		t.Fatalf("gallus != %s", s)
	}
}

func TestGenusGiraffa(t *testing.T) {
	var s string = "Giraffa"
	if ok := giraffa == s; !ok {
		t.Fatalf("giraffa != %s", s)
	}
}

func TestGenusGorilla(t *testing.T) {
	var s string = "Gorilla"
	if ok := gorilla == s; !ok {
		t.Fatalf("gorilla != %s", s)
	}
}

func TestGenusMacropus(t *testing.T) {
	var s string = "Macropus"
	if ok := macropus == s; !ok {
		t.Fatalf("macropus != %s", s)
	}
}

func TestGenusMus(t *testing.T) {
	var s string = "Mus"
	if ok := mus == s; !ok {
		t.Fatalf("mus != %s", s)
	}
}

func TestGenusNyctereutes(t *testing.T) {
	var s string = "Nyctereutes"
	if ok := nyctereutes == s; !ok {
		t.Fatalf("nyctereutes != %s", s)
	}
}

func TestGenusOdobenus(t *testing.T) {
	var s string = "Odobenus"
	if ok := odobenus == s; !ok {
		t.Fatalf("odobenus != %s", s)
	}
}

func TestGenusOvis(t *testing.T) {
	var s string = "Ovis"
	if ok := ovis == s; !ok {
		t.Fatalf("ovis != %s", s)
	}
}

func TestGenusPanthera(t *testing.T) {
	var s string = "Panthera"
	if ok := panthera == s; !ok {
		t.Fatalf("panthera != %s", s)
	}
}

func TestGenusPhascolarctos(t *testing.T) {
	var s string = "Phascolarctos"
	if ok := phascolarctos == s; !ok {
		t.Fatalf("phascolarctos != %s", s)
	}
}

func TestGenusPavo(t *testing.T) {
	var s string = "Pavo"
	if ok := pavo == s; !ok {
		t.Fatalf("pavo != %s", s)
	}
}

func TestGenusPelecanus(t *testing.T) {
	var s string = "Pelecanus"
	if ok := pelecanus == s; !ok {
		t.Fatalf("pelecanus != %s", s)
	}
}

func TestGenusRaphus(t *testing.T) {
	var s string = "Raphus"
	if ok := raphus == s; !ok {
		t.Fatalf("raphus != %s", s)
	}
}

func TestGenusRangifer(t *testing.T) {
	var s string = "Rangifer"
	if ok := rangifer == s; !ok {
		t.Fatalf("rangifer != %s", s)
	}
}

func TestGenusStruthio(t *testing.T) {
	var s string = "Struthio"
	if ok := struthio == s; !ok {
		t.Fatalf("struthio != %s", s)
	}
}

func TestGenusSus(t *testing.T) {
	var s string = "Sus"
	if ok := sus == s; !ok {
		t.Fatalf("sus != %s", s)
	}
}

func TestGenusTapirus(t *testing.T) {
	var s string = "Tapirus"
	if ok := tapirus == s; !ok {
		t.Fatalf("tapirus != %s", s)
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
