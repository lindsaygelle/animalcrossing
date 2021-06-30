package orders

import (
	"strings"
	"testing"
)

func TestAll(t *testing.T) {
	var (
		c = [...]string{
			accipitriformes,
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
