package astrology

import (
	"strings"
	"testing"
)

func TestL(t *testing.T) {
	for i := 0; i < len(L); i++ {
		var s string = strings.Title(L[i])
		if ok := L[i] == s; !ok {
			t.Fatalf("%s != %s; i=%d", L[i], s, i)
		}
		if ok := s[0] == 'L'; !ok {
			t.Fatalf("%s != %s; i=%d", string(s[0]), "L", i)
		}
	}
}
