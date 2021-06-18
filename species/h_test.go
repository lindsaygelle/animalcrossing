package species

import (
	"strings"
	"testing"
)

func TestH(t *testing.T) {
	for i := 0; i < len(H); i++ {
		var s string = strings.Title(H[i])
		if ok := H[i] == s; !ok {
			t.Fatalf("%s != %s; i=%d", H[i], s, i)
		}
		if ok := s[0] == 'H'; !ok {
			t.Fatalf("%s != %s; i=%d", string(s[0]), "H", i)
		}
	}
}
