package families

import (
	"strings"
	"testing"
)

func TestAll(t *testing.T) {
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
			myrmecophagidae,
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
