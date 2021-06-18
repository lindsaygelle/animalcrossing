package astrology

import (
	"strings"
	"testing"
)

func TestV(t *testing.T) {
	for i := 0; i < len(V); i++ {
		var s string = strings.Title(V[i])
		if ok := V[i] == s; !ok {
			t.Fatalf("%s != %s; i=%d", V[i], s, i)
		}
		if ok := s[0] == 'V'; !ok {
			t.Fatalf("%s != %s; i=%d", string(s[0]), "V", i)
		}
	}
}
