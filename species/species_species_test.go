package species

import (
	"strings"
	"testing"
)

func TestSpeciesAll(t *testing.T) {
	var (
		c = [...]string{aMexicanum,
			bTaurus,
			cAegagrus,
			canisLupusFamiliaris,
			cKingii,
			fCatus,
			gCamelopardalis,
			gGallus,
			nProcyonoides,
			oAries,
			oRosmarus,
			pCinereus,
			pLeo,
			rCucullatus,
			rTarandus,
			sCamelus,
			sScrofa,
			pTigris,
			vVulpes}
	)
	for i := 0; i < len(c); i++ {
		var s string = c[i]
		if ok := string(s[0]) == strings.ToUpper(string(s[0])); !ok {
			t.Fatalf("%s does not start with a capital letter", s)
		}
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

func TestSpeciesCAegagrus(t *testing.T) {
	var s string = "C. aegagrus"
	if ok := cAegagrus == s; !ok {
		t.Fatalf("cAegagrus != %s", s)
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

func TestSpeciesGCamelopardalis(t *testing.T) {
	var s string = "G. camelopardalis"
	if ok := gCamelopardalis == s; !ok {
		t.Fatalf("gCamelopardalis != %s", s)
	}
}

func TestSpeciesGGallus(t *testing.T) {
	var s string = "G. gallus"
	if ok := gGallus == s; !ok {
		t.Fatalf("gGallus != %s", s)
	}
}

func TestSpeciesOAries(t *testing.T) {
	var s string = "O. aries"
	if ok := oAries == s; !ok {
		t.Fatalf("oAries != %s", s)
	}
}

func TestSpeciesORosmarus(t *testing.T) {
	var s string = "O. rosmarus"
	if ok := oRosmarus == s; !ok {
		t.Fatalf("oRosmarus != %s", s)
	}
}

func TestSpeciesNProcyonoides(t *testing.T) {
	var s string = "N. procyonoides"
	if ok := nProcyonoides == s; !ok {
		t.Fatalf("nProcyonoides != %s", s)
	}
}

func TestSpeciesPCinereus(t *testing.T) {
	var s string = "P. cinereus"
	if ok := pCinereus == s; !ok {
		t.Fatalf("pCinereus != %s", s)
	}
}

func TestSpeciesPLeo(t *testing.T) {
	var s string = "P. leo"
	if ok := pLeo == s; !ok {
		t.Fatalf("pLeo != %s", s)
	}
}

func TestSpeciesRCucullatus(t *testing.T) {
	var s string = "R. cucullatus"
	if ok := rCucullatus == s; !ok {
		t.Fatalf("rCucullatus != %s", s)
	}
}

func TestSpeciesRTarandus(t *testing.T) {
	var s string = "R. tarandus"
	if ok := rTarandus == s; !ok {
		t.Fatalf("rTarandus != %s", s)
	}
}

func TestSpeciesSCamelus(t *testing.T) {
	var s string = "S. camelus"
	if ok := sCamelus == s; !ok {
		t.Fatalf("sCamelus != %s", s)
	}
}

func TestSpeciesSScrofa(t *testing.T) {
	var s string = "S. scrofa"
	if ok := sScrofa == s; !ok {
		t.Fatalf("sScrofa != %s", s)
	}
}

func TestSpeciesPTigris(t *testing.T) {
	var s string = "P. tigris"
	if ok := pTigris == s; !ok {
		t.Fatalf("pTigris != %s", s)
	}
}

func TestSpeciesVVulpes(t *testing.T) {
	var s string = "V. vulpes"
	if ok := vVulpes == s; !ok {
		t.Fatalf("vVulpes != %s", s)
	}
}
