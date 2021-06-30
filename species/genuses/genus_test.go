package genuses

import (
	"strings"
	"testing"
)

func TestAll(t *testing.T) {
	var (
		c = [...]string{
			alligator,
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
