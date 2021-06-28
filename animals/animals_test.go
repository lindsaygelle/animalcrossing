package animals

import (
	"strings"
	"testing"
)

func testAnimals(animals []Animal, letter string, t *testing.T) {
	for i := 0; i < len(animals); i++ {
		v := animals[i]
		a := v.Name()
		b := strings.Title(a)
		if ok := a == b; !ok {
			t.Fatalf("%s != %s; i=%d", a, b, i)
		}
		if ok := string(b[0]) == letter; !ok {
			t.Fatalf("%s != %s; i=%d", string(b[i]), letter, i)
		}
	}
}
