package animals

import (
	"strings"
	"testing"
)

func TestC(t *testing.T) {
	for i := 0; i < len(C); i++ {
		var s string = strings.Title(C[i])
		if ok := C[i] == s; !ok {
			t.Fatalf("%s != %s; i=%d", C[i], s, i)
		}
		if ok := s[0] == 'C'; !ok {
			t.Fatalf("%s != %s; i=%d", string(s[0]), "C", i)
		}
	}
}
